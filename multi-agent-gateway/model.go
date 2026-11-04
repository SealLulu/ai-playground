package multiagent

import (
	"context"
	"time"
)

// ModelClient is implemented by the model-specific REST adapter.
type ModelClient interface {
	Call(ctx context.Context, input AgentCallInput) AgentCallResult
}

// EchoModel is a small local model stand-in used by the runnable demo.
type EchoModel struct{}

func (EchoModel) Call(ctx context.Context, input AgentCallInput) AgentCallResult {
	select {
	case <-ctx.Done():
		return AgentCallResult{ErrorMessage: ctx.Err().Error()}
	default:
		return AgentCallResult{RESTCallSucceeded: true, Output: "Processed: " + input.TaskText}
	}
}

func callWithTimeout(parent context.Context, timeoutSeconds int, model ModelClient, input AgentCallInput) AgentCallResult {
	if timeoutSeconds <= 0 {
		return model.Call(parent, input)
	}
	ctx, cancel := context.WithTimeout(parent, time.Duration(timeoutSeconds)*time.Second)
	defer cancel()
	return model.Call(ctx, input)
}
