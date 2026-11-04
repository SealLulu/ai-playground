package multiagent

import (
	"errors"
	"strings"
	"unicode/utf8"
)

// Request is the public API input.
type Request struct {
	RequestID string `json:"request_id"`
	AgentID   string `json:"agent_id"`
	TaskText  string `json:"task_text"`
}

// APIResponse is the stable response returned to callers.
type APIResponse struct {
	RequestID    string `json:"request_id"`
	Status       int    `json:"status"`
	Output       string `json:"output,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}

type ValidationResult struct {
	Valid        bool
	ErrorMessage string
}

const (
	ResponseStatusSuccess  = 0
	ResponseStatusFailure  = 1
	ResponseStatusError    = -1
	DefaultTaskLengthLimit = 10000
)

var (
	ErrAgentNotFound   = errors.New("agent does not exist")
	ErrRequestConflict = errors.New("request ID is already used for different request data")
)

func ValidateRequest(request Request, maxTaskLength int) ValidationResult {
	switch {
	case strings.TrimSpace(request.RequestID) == "":
		return ValidationResult{ErrorMessage: "Request ID is required."}
	case strings.TrimSpace(request.AgentID) == "":
		return ValidationResult{ErrorMessage: "Agent ID is required."}
	case strings.TrimSpace(request.TaskText) == "":
		return ValidationResult{ErrorMessage: "Task text is required."}
	case maxTaskLength > 0 && utf8.RuneCountInString(request.TaskText) > maxTaskLength:
		return ValidationResult{ErrorMessage: "Task text exceeds the length limit."}
	default:
		return ValidationResult{Valid: true}
	}
}
