package multiagent

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

// AgentStatus is the availability state stored with an agent.
type AgentStatus string

const (
	AgentStatusAvailable   AgentStatus = "available"
	AgentStatusUnavailable AgentStatus = "unavailable"
)

// Agent contains the model routing and retry settings for one agent.
type Agent struct {
	AgentID        string
	ModelName      string
	Status         AgentStatus
	RetryLimit     int
	TimeoutSeconds int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// AgentCallInput is the task payload sent to a model.
type AgentCallInput struct {
	Agent    Agent
	TaskText string
}

// AgentCallResult is the result of one model REST call.
type AgentCallResult struct {
	RESTCallSucceeded bool
	Output            string
	ErrorMessage      string
}

// AgentRepository is backed by MySQL in production; MemoryAgentRepository is
// provided so this project can run locally without external services.
type AgentRepository interface {
	GetAgent(ctx context.Context, agentID string) (Agent, bool, error)
	SetAgentStatus(ctx context.Context, agentID string, status AgentStatus) error
}

// MemoryAgentRepository is a concurrency-safe local agent repository.
type MemoryAgentRepository struct {
	mu     sync.RWMutex
	agents map[string]Agent
}

func NewMemoryAgentRepository(agents ...Agent) *MemoryAgentRepository {
	repo := &MemoryAgentRepository{agents: make(map[string]Agent, len(agents))}
	for _, agent := range agents {
		repo.agents[agent.AgentID] = agent
	}
	return repo
}

func (r *MemoryAgentRepository) GetAgent(_ context.Context, agentID string) (Agent, bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	agent, ok := r.agents[agentID]
	return agent, ok, nil
}

func (r *MemoryAgentRepository) SetAgentStatus(_ context.Context, agentID string, status AgentStatus) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	agent, ok := r.agents[agentID]
	if !ok {
		return ErrAgentNotFound
	}
	agent.Status = status
	agent.UpdatedAt = time.Now().UTC()
	r.agents[agentID] = agent
	return nil
}

// FailureCounters stores consecutive REST failures by agent ID.
// Atomic counters and sync.Map make updates safe across concurrent requests.
type FailureCounters struct {
	counts sync.Map // map[string]*atomic.Int64
}

func (f *FailureCounters) Increment(agentID string) int64 {
	counter, _ := f.counts.LoadOrStore(agentID, &atomic.Int64{})
	return counter.(*atomic.Int64).Add(1)
}

func (f *FailureCounters) Delete(agentID string) {
	f.counts.Delete(agentID)
}

func (f *FailureCounters) Get(agentID string) (int64, bool) {
	value, ok := f.counts.Load(agentID)
	if !ok {
		return 0, false
	}
	return value.(*atomic.Int64).Load(), true
}
