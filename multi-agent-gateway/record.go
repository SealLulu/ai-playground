package multiagent

import (
	"context"
	"sync"
	"time"
)

const (
	ProcessingStatusRunning   = "processing"
	ProcessingStatusSucceeded = "succeeded"
	ProcessingStatusFailed    = "failed"
)

type RequestRecord struct {
	Request   Request
	CreatedAt time.Time
}

type ProcessingRecord struct {
	RequestID    string
	Status       string
	AttemptCount int
	FailureCount int
	Result       string
	ErrorDetails string
	StartedAt    time.Time
	CompletedAt  time.Time
}

// RecordStore is backed by Redis in production. Claim provides the atomic
// request-ID lock used to coalesce concurrent duplicate requests.
type RecordStore interface {
	SaveRequest(context.Context, RequestRecord) error
	SaveProcessing(context.Context, ProcessingRecord) error
	GetProcessing(context.Context, string) (ProcessingRecord, bool, error)
	GetResponse(context.Context, string) (APIResponse, bool, error)
	Claim(context.Context, string) (owner bool, done <-chan struct{}, err error)
	Complete(context.Context, string, APIResponse) error
}

type requestFlight struct {
	done     chan struct{}
	response APIResponse
}

// MemoryRecordStore is a concurrency-safe local stand-in for Redis.
type MemoryRecordStore struct {
	mu         sync.RWMutex
	requests   map[string]RequestRecord
	processing map[string]ProcessingRecord
	responses  map[string]APIResponse
	inFlight   map[string]*requestFlight
}

func NewMemoryRecordStore() *MemoryRecordStore {
	return &MemoryRecordStore{
		requests: make(map[string]RequestRecord), processing: make(map[string]ProcessingRecord),
		responses: make(map[string]APIResponse), inFlight: make(map[string]*requestFlight),
	}
}

func (s *MemoryRecordStore) SaveRequest(_ context.Context, record RequestRecord) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.requests[record.Request.RequestID] = record
	return nil
}

func (s *MemoryRecordStore) SaveProcessing(_ context.Context, record ProcessingRecord) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.processing[record.RequestID] = record
	return nil
}

func (s *MemoryRecordStore) GetProcessing(_ context.Context, requestID string) (ProcessingRecord, bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	record, ok := s.processing[requestID]
	return record, ok, nil
}

func (s *MemoryRecordStore) GetResponse(_ context.Context, requestID string) (APIResponse, bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	response, ok := s.responses[requestID]
	return response, ok, nil
}

func (s *MemoryRecordStore) Claim(_ context.Context, requestID string) (bool, <-chan struct{}, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, completed := s.responses[requestID]; completed {
		done := make(chan struct{})
		close(done)
		return false, done, nil
	}
	if flight, exists := s.inFlight[requestID]; exists {
		return false, flight.done, nil
	}
	flight := &requestFlight{done: make(chan struct{})}
	s.inFlight[requestID] = flight
	return true, flight.done, nil
}

func (s *MemoryRecordStore) Complete(_ context.Context, requestID string, response APIResponse) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.responses[requestID] = response
	if flight, ok := s.inFlight[requestID]; ok {
		flight.response = response
		delete(s.inFlight, requestID)
		close(flight.done)
	}
	return nil
}

// Snapshot helpers are used by tests and local inspection.
func (s *MemoryRecordStore) Request(requestID string) (RequestRecord, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	record, ok := s.requests[requestID]
	return record, ok
}
