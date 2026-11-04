package multiagent

import (
	"context"
	"sync"
	"testing"
	"time"
)

type scriptedModel struct {
	mu      sync.Mutex
	results []AgentCallResult
	calls   int
	delay   time.Duration
}

func (m *scriptedModel) Call(_ context.Context, _ AgentCallInput) AgentCallResult {
	if m.delay > 0 {
		time.Sleep(m.delay)
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	m.calls++
	index := m.calls - 1
	if index >= len(m.results) {
		return AgentCallResult{RESTCallSucceeded: true, Output: "ok"}
	}
	return m.results[index]
}

func (m *scriptedModel) CallCount() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.calls
}

type testHarness struct {
	processor *Processor
	agents    *MemoryAgentRepository
	records   *MemoryRecordStore
	model     *scriptedModel
}

func newHarness(retryLimit int, results ...AgentCallResult) *testHarness {
	agents := NewMemoryAgentRepository(Agent{
		AgentID: "agent-001", ModelName: "test-model", Status: AgentStatusAvailable,
		RetryLimit: retryLimit, TimeoutSeconds: 1,
	})
	records := NewMemoryRecordStore()
	model := &scriptedModel{results: results}
	return &testHarness{NewProcessor(agents, records, model), agents, records, model}
}

func TestValidRequestSucceeds(t *testing.T) {
	h := newHarness(2, AgentCallResult{RESTCallSucceeded: true, Output: "The note says Friday."})
	got := h.processor.ProcessRequest(context.Background(), Request{"req-001", "agent-001", "Summarize this note."})
	if got.Status != ResponseStatusSuccess || got.Output != "The note says Friday." {
		t.Fatalf("unexpected response: %+v", got)
	}
	record, found, _ := h.records.GetProcessing(context.Background(), "req-001")
	if !found || record.Status != ProcessingStatusSucceeded || record.AttemptCount != 1 {
		t.Fatalf("unexpected processing record: %+v, found=%v", record, found)
	}
}

func TestUnknownAgentID(t *testing.T) {
	h := newHarness(1)
	got := h.processor.ProcessRequest(context.Background(), Request{"req-002", "agent-missing", "task"})
	if got.Status != ResponseStatusFailure || got.ErrorMessage != "Agent does not exist." || h.model.CallCount() != 0 {
		t.Fatalf("unexpected response/calls: %+v, calls=%d", got, h.model.CallCount())
	}
}

func TestUnavailableAgent(t *testing.T) {
	h := newHarness(1)
	if err := h.agents.SetAgentStatus(context.Background(), "agent-001", AgentStatusUnavailable); err != nil {
		t.Fatal(err)
	}
	got := h.processor.ProcessRequest(context.Background(), Request{"req-003", "agent-001", "task"})
	if got.Status != ResponseStatusFailure || got.ErrorMessage != "Agent is unavailable." || h.model.CallCount() != 0 {
		t.Fatalf("unexpected response/calls: %+v, calls=%d", got, h.model.CallCount())
	}
}

func TestEmptyTaskText(t *testing.T) {
	h := newHarness(1)
	got := h.processor.ProcessRequest(context.Background(), Request{"req-004", "agent-001", "   "})
	if got.Status != ResponseStatusFailure || h.model.CallCount() != 0 {
		t.Fatalf("expected validation rejection without model call, got %+v", got)
	}
}

func TestTaskLengthBoundary(t *testing.T) {
	h := newHarness(0, AgentCallResult{RESTCallSucceeded: true, Output: "ok"})
	atLimit := Request{"req-005a", "agent-001", string(make([]byte, DefaultTaskLengthLimit))}
	if got := h.processor.ProcessRequest(context.Background(), atLimit); got.Status != ResponseStatusSuccess {
		t.Fatalf("task at limit should pass: %+v", got)
	}
	overLimit := Request{"req-005b", "agent-001", string(make([]byte, DefaultTaskLengthLimit+1))}
	got := h.processor.ProcessRequest(context.Background(), overLimit)
	if got.Status != ResponseStatusFailure || h.model.CallCount() != 1 {
		t.Fatalf("task over limit should be rejected before model call: %+v calls=%d", got, h.model.CallCount())
	}
}

func TestConcurrentDuplicateRequestSharesResult(t *testing.T) {
	h := newHarness(1, AgentCallResult{RESTCallSucceeded: true, Output: "shared"})
	h.model.delay = 30 * time.Millisecond
	req := Request{"req-006", "agent-001", "same task"}
	const count = 8
	responses := make([]APIResponse, count)
	var wg sync.WaitGroup
	for i := range responses {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			responses[i] = h.processor.ProcessRequest(context.Background(), req)
		}(i)
	}
	wg.Wait()
	for _, response := range responses {
		if response.Status != ResponseStatusSuccess || response.Output != "shared" {
			t.Fatalf("duplicate did not receive shared result: %+v", response)
		}
	}
	if calls := h.model.CallCount(); calls != 1 {
		t.Fatalf("expected one model call, got %d", calls)
	}
}

func TestModelCallRetryLimitMarksAgentUnavailable(t *testing.T) {
	timeout := AgentCallResult{ErrorMessage: "model call timed out"}
	h := newHarness(2, timeout, timeout, timeout)
	got := h.processor.ProcessRequest(context.Background(), Request{"req-007", "agent-001", "task"})
	if got.Status != ResponseStatusError || h.model.CallCount() != 3 {
		t.Fatalf("expected failure after three attempts, got %+v calls=%d", got, h.model.CallCount())
	}
	record, found, _ := h.records.GetProcessing(context.Background(), "req-007")
	agent, _, _ := h.agents.GetAgent(context.Background(), "agent-001")
	failures, exists := h.processor.Failures.Get("agent-001")
	if !found || record.AttemptCount != 3 || record.FailureCount != 3 || !exists || failures != 3 || agent.Status != AgentStatusUnavailable {
		t.Fatalf("retry outcome mismatch: record=%+v failure=%d exists=%v agent=%+v", record, failures, exists, agent)
	}
}

func TestSuccessfulRESTCallClearsFailureCounter(t *testing.T) {
	h := newHarness(1, AgentCallResult{RESTCallSucceeded: true, Output: "recovered"})
	h.processor.Failures.Increment("agent-001")
	got := h.processor.ProcessRequest(context.Background(), Request{"req-008", "agent-001", "task"})
	if got.Status != ResponseStatusSuccess {
		t.Fatalf("expected success: %+v", got)
	}
	if _, exists := h.processor.Failures.Get("agent-001"); exists {
		t.Fatal("failure counter key was not cleared")
	}
}

func TestInvalidModelOutputRetriesAndMarksAgentUnavailable(t *testing.T) {
	invalid := AgentCallResult{RESTCallSucceeded: true, Output: ""}
	h := newHarness(1, invalid, invalid)
	got := h.processor.ProcessRequest(context.Background(), Request{"req-009", "agent-001", "task"})
	record, found, _ := h.records.GetProcessing(context.Background(), "req-009")
	agent, _, _ := h.agents.GetAgent(context.Background(), "agent-001")
	if got.Status != ResponseStatusError || h.model.CallCount() != 2 || !found ||
		record.Status != ProcessingStatusFailed || record.AttemptCount != 2 ||
		record.ErrorDetails == "" || agent.Status != AgentStatusUnavailable {
		t.Fatalf("invalid output retry mismatch: response=%+v record=%+v agent=%+v", got, record, agent)
	}
}

func TestRequestAndProcessingRecordsPersist(t *testing.T) {
	h := newHarness(0, AgentCallResult{RESTCallSucceeded: true, Output: "saved"})
	request := Request{"req-010", "agent-001", "task without secrets"}
	got := h.processor.ProcessRequest(context.Background(), request)
	requestRecord, requestFound := h.records.Request(request.RequestID)
	processing, processingFound, _ := h.records.GetProcessing(context.Background(), request.RequestID)
	if got.Status != ResponseStatusSuccess || !requestFound || !processingFound ||
		requestRecord.Request != request || processing.Result != "saved" ||
		processing.Status != ProcessingStatusSucceeded {
		t.Fatalf("records missing or inconsistent: response=%+v request=%+v processing=%+v", got, requestRecord, processing)
	}
}
