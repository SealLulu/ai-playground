package multiagent

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type Processor struct {
	Agents        AgentRepository
	Records       RecordStore
	Model         ModelClient
	Failures      *FailureCounters
	MaxTaskLength int
}

func NewProcessor(agents AgentRepository, records RecordStore, model ModelClient) *Processor {
	return &Processor{
		Agents: agents, Records: records, Model: model,
		Failures: &FailureCounters{}, MaxTaskLength: DefaultTaskLengthLimit,
	}
}

func (p *Processor) ProcessRequest(ctx context.Context, request Request) APIResponse {
	validation := ValidateRequest(request, p.MaxTaskLength)
	if !validation.Valid {
		return APIResponse{RequestID: request.RequestID, Status: ResponseStatusFailure, ErrorMessage: validation.ErrorMessage}
	}
	if p.Agents == nil || p.Records == nil || p.Model == nil {
		return APIResponse{RequestID: request.RequestID, Status: ResponseStatusError, ErrorMessage: "Processing service is not configured."}
	}

	if response, found, err := p.Records.GetResponse(ctx, request.RequestID); err == nil && found {
		return response
	}
	owner, done, err := p.Records.Claim(ctx, request.RequestID)
	if err != nil {
		return APIResponse{RequestID: request.RequestID, Status: ResponseStatusError, ErrorMessage: err.Error()}
	}
	if !owner {
		select {
		case <-done:
			if response, found, err := p.Records.GetResponse(ctx, request.RequestID); err == nil && found {
				return response
			}
			return APIResponse{RequestID: request.RequestID, Status: ResponseStatusError, ErrorMessage: "Could not read the shared request result."}
		case <-ctx.Done():
			return APIResponse{RequestID: request.RequestID, Status: ResponseStatusError, ErrorMessage: ctx.Err().Error()}
		}
	}

	response := p.processOwner(ctx, request)
	if err := p.Records.Complete(ctx, request.RequestID, response); err != nil && response.Status == ResponseStatusSuccess {
		return APIResponse{RequestID: request.RequestID, Status: ResponseStatusError, ErrorMessage: "Could not save the request result."}
	}
	return response
}

func (p *Processor) processOwner(ctx context.Context, request Request) APIResponse {
	agent, found, err := p.Agents.GetAgent(ctx, request.AgentID)
	if err != nil {
		return APIResponse{RequestID: request.RequestID, Status: ResponseStatusError, ErrorMessage: err.Error()}
	}
	if !found {
		return APIResponse{RequestID: request.RequestID, Status: ResponseStatusFailure, ErrorMessage: "Agent does not exist."}
	}
	if agent.Status != AgentStatusAvailable {
		return APIResponse{RequestID: request.RequestID, Status: ResponseStatusFailure, ErrorMessage: "Agent is unavailable."}
	}

	now := time.Now().UTC()
	if err := p.Records.SaveRequest(ctx, RequestRecord{Request: request, CreatedAt: now}); err != nil {
		return APIResponse{RequestID: request.RequestID, Status: ResponseStatusError, ErrorMessage: "Could not save request."}
	}
	record := ProcessingRecord{RequestID: request.RequestID, Status: ProcessingStatusRunning, StartedAt: now}
	if err := p.Records.SaveProcessing(ctx, record); err != nil {
		return APIResponse{RequestID: request.RequestID, Status: ResponseStatusError, ErrorMessage: "Could not start processing record."}
	}

	var lastError string
	for attempt := 0; attempt <= agent.RetryLimit; attempt++ {
		record.AttemptCount = attempt + 1
		result := callWithTimeout(ctx, agent.TimeoutSeconds, p.Model, AgentCallInput{Agent: agent, TaskText: request.TaskText})
		if result.RESTCallSucceeded {
			p.Failures.Delete(agent.AgentID)
			if strings.TrimSpace(result.Output) != "" {
				record.Status = ProcessingStatusSucceeded
				record.Result = result.Output
				record.CompletedAt = time.Now().UTC()
				_ = p.Records.SaveProcessing(ctx, record)
				return APIResponse{RequestID: request.RequestID, Status: ResponseStatusSuccess, Output: result.Output}
			}
			lastError = "Agent returned empty output."
			record.FailureCount++
			if attempt == agent.RetryLimit {
				_ = p.Agents.SetAgentStatus(ctx, agent.AgentID, AgentStatusUnavailable)
			}
			continue
		}

		lastError = result.ErrorMessage
		if lastError == "" {
			lastError = "Agent REST call failed."
		}
		record.FailureCount++
		if p.Failures.Increment(agent.AgentID) >= int64(agent.RetryLimit+1) {
			_ = p.Agents.SetAgentStatus(ctx, agent.AgentID, AgentStatusUnavailable)
			break
		}
	}

	record.Status = ProcessingStatusFailed
	record.ErrorDetails = lastError
	record.CompletedAt = time.Now().UTC()
	_ = p.Records.SaveProcessing(ctx, record)
	return APIResponse{RequestID: request.RequestID, Status: ResponseStatusError, ErrorMessage: fmt.Sprintf("Agent request failed: %s", lastError)}
}
