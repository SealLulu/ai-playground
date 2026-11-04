# Multi-Agent — System Design
# Multi-Agent — 系统设计

> **Language rule:** Write every document section in English, with its Simplified Chinese translation immediately after each English sentence.
> **语言规则：** 文档的每个部分都先用英文书写，并在每个英文句子后紧接简体中文翻译。

> **Writing style:** Use simple, clear wording.
> **写作风格：** 使用简单、清晰的表达。

## 1. Summary
## 1. 概述

### 1.1 Main Functions
### 1.1 主要功能

#### 1.1.1 Handle Requests
#### 1.1.1 处理请求

Each request contains an agent ID and a task.
每个请求都包含智能体 ID 和任务。

The system uses the agent ID to route the task to the model assigned to that agent.
系统根据智能体 ID，将任务路由到分配给该智能体的模型。

The system calls the selected model and returns its response to the user.
系统调用选定的模型，并将其答复返回给用户。

#### 1.1.2 Track Agent Availability
#### 1.1.2 跟踪智能体可用状态

The system tracks whether each agent is available to handle a request.
系统跟踪每个智能体是否可以处理请求。

An agent's status can change between available and unavailable.
智能体的状态可以在可用和不可用之间切换。

### 1.2 Edge Cases
### 1.2 边界情况

#### 1.2.1 Agent Is Unavailable
#### 1.2.1 智能体不可用

If the agent ID exists but the agent is unavailable, the system does not call its model.
如果智能体 ID 存在但该智能体不可用，系统不会调用它的模型。

The system tells the user that the agent is unavailable and asks them to try again later.
系统会告知用户该智能体当前不可用，并请用户稍后重试。

#### 1.2.2 Agent ID Does Not Exist
#### 1.2.2 智能体 ID 不存在

If the agent ID does not exist, the system does not call a model.
如果智能体 ID 不存在，系统不会调用模型。

The system returns a clear error asking the user to provide a valid agent ID.
系统会返回清晰的错误信息，请用户提供有效的智能体 ID。

#### 1.2.3 Request Text Is Empty
#### 1.2.3 请求文本为空

If the task or chat text is empty, the system does not call a model.
如果任务或聊天文本为空，系统不会调用模型。

The system asks the user to provide the task text.
系统会请用户提供任务文本。

#### 1.2.4 Request Is Repeated
#### 1.2.4 请求重复提交

If two copies of the same request arrive at the same time with the same request ID, the system runs the task once and shares the result with both requests.
如果同一个请求的两个副本同时使用相同的请求 ID 到达，系统只执行一次任务，并将结果返回给两个请求。

The system needs a request ID and stored request status to detect and handle repeated requests safely.
系统需要请求 ID 和已保存的请求状态，才能安全地识别并处理重复请求。

#### 1.2.5 Model Fails Repeatedly
#### 1.2.5 模型多次调用失败

If the model call fails repeatedly and reaches the retry limit, the system marks the agent as unavailable.
如果模型调用多次失败并达到重试上限，系统会将该智能体标记为不可用。

The system returns a clear failure message to the user.
系统会向用户返回清晰的失败信息。

#### 1.2.6 Request Is Too Long
#### 1.2.6 请求过长

If the task exceeds the allowed length, the system does not call a model.
如果任务超过允许的长度，系统不会调用模型。

The system asks the user to shorten the task.
系统会请用户缩短任务内容。

Multi-Agent is a software project that coordinates multiple agents to complete tasks.
Multi-Agent 是一个协调多个智能体共同完成任务的软件项目。

The system receives a goal, assigns work, and combines the results.
系统接收一个目标，分配工作，并整合结果。

This document describes the main system design at a high level.
本文档从高层次介绍系统的主要设计。

## 2. Business Logic
## 2. 业务逻辑

### 2.1 Main System Items
### 2.1 系统中的主要项目

#### 2.1.1 Agent
#### 2.1.1 智能体

##### 2.1.1.1 Change Agent Status
##### 2.1.1.1 更改智能体状态

The system tracks whether each agent is available or unavailable.
系统跟踪每个智能体处于可用还是不可用状态。

If calls to an agent repeatedly fail and reach the retry limit, the system changes its status to unavailable.
如果对智能体的调用多次失败并达到重试上限，系统会将其状态更改为不可用。

When an agent REST call succeeds, the system removes that agent's failure-count entry from the in-memory map.
智能体 REST 调用成功时，系统会从内存映射中删除该智能体的失败计数项。

The agent can be marked available again when it is ready to handle requests.
智能体准备好处理请求后，可以再次将其标记为可用。

##### 2.1.1.2 Route a Request and Return the Response
##### 2.1.1.2 路由请求并返回响应

When a new request arrives, the system uses its agent ID to find and route the request to that agent.
收到新请求时，系统会根据请求中的智能体 ID 查找对应智能体，并将请求路由给它。

The system calls the agent and sends its response back to the user.
系统调用智能体，并将智能体的响应返回给用户。

##### 2.1.1.3 Call the Agent and Check the Response
##### 2.1.1.3 调用智能体并检查响应

The system prepares the input, sends it to the Agent API, receives the response, and checks the chat output.
系统准备输入，将其发送到智能体 API，接收响应，并检查聊天输出。

If the API call fails or the chat output is invalid, the system retries up to the retry limit.
如果 API 调用失败或聊天输出无效，系统会在重试上限内重试。

If the retry limit is reached, the system records the failure and marks the agent unavailable.
如果达到重试上限，系统会记录失败并将智能体标记为不可用。

If the chat output is valid, the system records success and sends the output to the user.
如果聊天输出有效，系统会记录成功并将输出发送给用户。

See the editable [Agent-Call-Workflow.drawio](Agent-Call-Workflow.drawio) diagram for this process.
此流程请参见可编辑的 [Agent-Call-Workflow.drawio](Agent-Call-Workflow.drawio) 图。

#### 2.1.2 Request
#### 2.1.2 请求

A request has a request ID, an agent ID, and task text.
每个请求都包含请求 ID、智能体 ID 和任务文本。

The system uses the request ID to recognize repeated requests.
系统使用请求 ID 来识别重复请求。

##### 2.1.2.1 Validate the Request
##### 2.1.2.1 检查请求

The system checks that the request has a valid agent ID and non-empty task text.
系统会检查请求是否包含有效的智能体 ID，以及非空的任务文本。

The system rejects task text that exceeds the allowed length and does not call the agent.
如果任务文本超过允许的长度，系统会拒绝该请求且不调用智能体。

For an invalid request, the system sends a clear error response to the caller.
如果请求无效，系统会向调用方返回清晰的错误信息。

##### 2.1.2.2 Map the API Request and Response
##### 2.1.2.2 转换 API 请求和响应

For a valid request, the system maps the API request fields into the input format required by the agent's model.
对于有效请求，系统会将 API 请求字段转换为智能体模型所需的输入格式。

The system maps the agent's output into a consistent API response and sends it to the caller.
系统会将智能体的输出转换为统一的 API 响应格式，并发送给调用方。

#### 2.1.3 Processing Record
#### 2.1.3 处理记录

A processing record tracks the request status, model call attempts, and result.
处理记录跟踪请求状态、模型调用次数和处理结果。

##### 2.1.3.1 Concurrent Duplicate Requests
##### 2.1.3.1 并发重复请求

The system creates one processing record for each request ID, and the request ID must be unique.
系统为每个请求 ID 创建一条处理记录，并要求请求 ID 唯一。

If several copies with the same request ID arrive at the same time, only one copy starts the model call.
如果多个使用相同请求 ID 的副本同时到达，系统只让其中一个副本发起模型调用。

The other copies use the same processing record and receive its status and final result.
其他副本使用同一条处理记录，并获取该请求的状态和最终结果。

##### 2.1.3.2 Model Call Failures
##### 2.1.3.2 模型调用失败

For each failed model call, the system records the attempt count and error details in the processing record.
每次模型调用失败时，系统都会在处理记录中保存尝试次数和错误详情。

The system retries up to the retry limit and then records the request as failed.
系统会在重试上限内重试，然后将请求记录为失败。

When the retry limit is reached, the system also marks the agent as unavailable.
达到重试上限后，系统还会将该智能体标记为不可用。

### 2.2 Workflow
### 2.2 工作流程

The main processing flow is shown in [Main-Processing-Workflow.drawio](Main-Processing-Workflow.drawio).
主要处理流程请参见 [Main-Processing-Workflow.drawio](Main-Processing-Workflow.drawio)。

## 3. Storage
## 3. 存储

### 3.1 Main Storage
### 3.1 主要存储

MySQL stores the authoritative agent records, including each agent's status and model settings.
MySQL 保存权威的智能体记录，包括每个智能体的状态和模型设置。

Redis stores request and processing records.
Redis 保存请求记录和处理记录。

The system keeps per-agent failure counters in a thread-safe in-memory map.
系统使用线程安全的内存映射保存每个智能体的失败计数。

Each Redis record uses a key and a value or a hash of named field-value pairs.
每条 Redis 记录都使用一个键和一个值，或使用由命名字段和值组成的哈希。

The system claims a request with an atomic `SET request-lock:<request_id> processing NX EX 60` operation so concurrent copies do not start duplicate processing.
系统使用原子 `SET request-lock:<request_id> processing NX EX 60` 操作占用请求，避免并发副本重复启动处理。

MySQL and Redis persistence and backups are enabled so stored records can be recovered after a restart.
系统会启用 MySQL 和 Redis 的持久化及备份，以便在重启后恢复已保存的记录。

### 3.2 Storage Structure
### 3.2 存储结构

#### 3.2.1 Agent Structure
#### 3.2.1 智能体结构

This Go example shows the agent structure and its status constants.
以下 Go 示例展示智能体结构及其状态常量。

```go
import "time"

type AgentStatus string

const (
	AgentStatusAvailable   AgentStatus = "available"
	AgentStatusUnavailable AgentStatus = "unavailable"
)

type Agent struct {
	AgentID       string
	ModelName     string
	Status        AgentStatus
	RetryLimit    int
	TimeoutSeconds int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
```

These Go structures define the input sent to the agent and the result returned by its REST call.
以下 Go 结构定义发送给智能体的输入，以及其 REST 调用返回的结果。

```go
type AgentCallInput struct {
	Agent    Agent
	TaskText string
}

type AgentCallResult struct {
	RESTCallSucceeded bool
	Output            string
	ErrorMessage      string
}
```

`retry_limit` sets the maximum number of retries after the first failed call.
`retry_limit` 用于设置首次调用失败后的最大重试次数。

`timeout_seconds` sets the maximum wait time for one model call; `30` is an example value.
`timeout_seconds` 用于设置单次模型调用的最长等待时间；`30` 是示例值。

#### 3.2.2 Request and Response Structures
#### 3.2.2 请求和响应结构

This JSON shows the fields the caller sends in a request.
以下 JSON 展示调用方在请求中发送的字段。

```json
{
  "request_id": "req-001",
  "agent_id": "agent-001",
  "task_text": "Complete the requested task."
}
```

This JSON shows the response returned to the caller.
以下 JSON 展示返回给调用方的响应。

```json
{
  "request_id": "req-001",
  "status": 0,
  "output": "Task result.",
  "error_message": null
}
```

The response status is `0` for success, `1` for failure, or `-1` for failure caused by an error.
响应状态为 `0` 表示成功，`1` 表示失败，`-1` 表示因错误而失败。

For a failed request, `output` is `null` and `error_message` contains the error message.
请求失败时，`output` 为 `null`，`error_message` 中包含错误信息。

These Go structures define the API request, API response, and validation result.
以下 Go 结构定义 API 请求、API 响应和验证结果。

```go
type Request struct {
	RequestID string
	AgentID   string
	TaskText  string
}

type APIResponse struct {
	RequestID    string
	Status       int
	Output       string
	ErrorMessage string
}

type ValidationResult struct {
	Valid        bool
	ErrorMessage string
}

const (
	ResponseStatusSuccess = 0
	ResponseStatusFailure = 1
	ResponseStatusError   = -1
)
```

#### 3.2.3 Record Structures
#### 3.2.3 记录结构

##### 3.2.3.1 Agent Record in MySQL
##### 3.2.3.1 MySQL 中的智能体记录

This SQL structure stores the authoritative agent settings and status in MySQL.
以下 SQL 结构在 MySQL 中保存权威的智能体设置和状态。

```sql
CREATE TABLE agents (
    agent_id VARCHAR(64) PRIMARY KEY,
    model_name VARCHAR(128) NOT NULL,
    status VARCHAR(16) NOT NULL,
    retry_limit INT NOT NULL,
    timeout_seconds INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
```

The `status` column stores the value from `AgentStatusAvailable` or `AgentStatusUnavailable`.
`status` 列保存 `AgentStatusAvailable` 或 `AgentStatusUnavailable` 对应的值。

##### 3.2.3.2 In-Memory Agent Failure Map
##### 3.2.3.2 内存中的智能体失败映射

The map uses the agent ID as the key and the consecutive failure count as the value.
该映射使用智能体 ID 作为键，并使用连续失败次数作为值。

```go
import (
	"sync"
	"sync/atomic"
)

var failureCounts sync.Map // key: agent ID; value: *atomic.Int64

// On a failed REST call, increment the count for this agent.
// REST 调用失败时，增加该智能体的失败计数。
counter, _ := failureCounts.LoadOrStore(agentID, &atomic.Int64{})
failureCount := counter.(*atomic.Int64).Add(1)

// On a successful REST call, remove the agent's failure-count key.
// REST 调用成功时，删除该智能体的失败计数键。
failureCounts.Delete(agentID)
```

The Go `sync.Map` and atomic counter update each agent's count safely when calls run at the same time.
当多个调用同时运行时，Go 的 `sync.Map` 和原子计数器可以安全地更新每个智能体的计数。

When the failure count reaches the initial call plus the configured retry limit, the system updates the agent status to unavailable in MySQL.
当失败次数达到首次调用次数加配置的重试上限时，系统会在 MySQL 中将智能体状态更新为不可用。

The in-memory failure counts are cleared when the service restarts.
服务重启时，内存中的失败计数会被清除。

##### 3.2.3.3 Request Record in Redis
##### 3.2.3.3 Redis 中的请求记录

This Redis command stores the request fields under a key formed from its request ID.
以下 Redis 命令将请求字段保存在由请求 ID 组成的键中。

```redis
HSET request:req-001 request_id req-001 agent_id agent-001 task_text "Complete the requested task." created_at 2026-11-03T10:00:00Z
```

This Go structure defines the request record stored in Redis.
以下 Go 结构定义保存在 Redis 中的请求记录。

```go
import "time"

type RequestRecord struct {
	RequestID string
	AgentID   string
	TaskText  string
	CreatedAt time.Time
}
```

##### 3.2.3.4 Processing Record in Redis
##### 3.2.3.4 Redis 中的处理记录

This Redis command stores the processing status, attempt count, failure count, result, and latest error for a request.
以下 Redis 命令为请求保存处理状态、尝试次数、失败次数、结果和最近一次错误。

```redis
HSET processing:req-001 request_id req-001 status succeeded attempt_count 1 failure_count 0 result "Task result." error_details "" started_at 2026-11-03T10:00:01Z completed_at 2026-11-03T10:00:05Z
```

This Go structure defines the processing record stored for a request.
以下 Go 结构定义为请求保存的处理记录。

```go
import "time"

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
```

`failure_count` counts failed agent calls for this request, and it increases after each failed call.
`failure_count` 统计该请求中智能体调用失败的次数，并在每次调用失败后递增。

The system can increment this value with `HINCRBY processing:req-001 failure_count 1`.
系统可以使用 `HINCRBY processing:req-001 failure_count 1` 增加该值。

The system stores project settings, task details, agent assignments, and task status.
系统保存项目设置、任务详情、智能体分配信息和任务状态。

The system stores processing records in persistent storage, using a unique request ID to prevent concurrent duplicate processing.
系统将处理记录保存在持久化存储中，并使用唯一请求 ID 防止并发重复处理。

Each processing record stores the current status, attempt count, failure details, and final result.
每条处理记录都保存当前状态、尝试次数、失败详情和最终结果。

It may store agent outputs and task history to support review and recovery.
系统可以保存智能体输出和任务历史，以便进行审查和恢复。

Sensitive data should be protected and kept only as long as needed.
敏感数据应受到保护，并且只在必要时保存。

## 4. Main Processing Code
## 4. 主要处理代码

### 4.1 Main Request Processing
### 4.1 请求主处理流程

The following Go-style outline shows the main request flow, with comments for each stage and calls to key helper functions.
以下 Go 风格的流程框架通过分段注释展示请求主流程，并调用关键辅助函数。

```go
func ProcessRequest(request Request) APIResponse {
	// Validate the request and return a failure response when it is invalid.
	// 验证请求；如果请求无效，则返回失败响应。
	validation := ValidateRequest(request)
	if !validation.Valid {
		return APIResponse{RequestID: request.RequestID, Status: ResponseStatusFailure, ErrorMessage: validation.ErrorMessage}
	}

	// Reuse an existing result or acquire the lock for this request ID.
	// 复用已有结果，或获取该请求 ID 的处理锁。
	if response, found, _ := GetSavedResponse(request.RequestID); found {
		return response
	}
	acquired, _ := TryAcquireRequestLock(request.RequestID)
	if !acquired {
		return WaitForRequestResult(request.RequestID)
	}
	defer ReleaseRequestLock(request.RequestID)

	// Load the agent and stop if it does not exist or is unavailable.
	// 加载智能体；如果智能体不存在或不可用，则停止处理。
	agent, found, _ := LoadAgent(request.AgentID)
	if !found || agent.Status != AgentStatusAvailable {
		return APIResponse{RequestID: request.RequestID, Status: ResponseStatusFailure, ErrorMessage: "Agent is unavailable or does not exist."}
	}

	// Save the request and start a processing record.
	// 保存请求并创建处理记录。
	_ = SaveRequestRecord(request)
	record := ProcessingRecord{RequestID: request.RequestID, Status: "processing"}
	_ = SaveProcessingRecord(record)

	// Make the first REST call and the configured number of retries.
	// 执行首次 REST 调用及配置的重试次数。
	for attempt := 0; attempt <= agent.RetryLimit; attempt++ {
		record.AttemptCount = attempt + 1
		result := CallAgent(AgentCallInput{Agent: agent, TaskText: request.TaskText}, agent.TimeoutSeconds)

		// A successful REST call clears the agent's failure-count key.
		// REST 调用成功后，清除智能体的失败计数键。
		if result.RESTCallSucceeded {
			ClearAgentFailureCount(agent.AgentID)
			if ValidateAgentOutput(result.Output).Valid {
				record.Status = "succeeded"
				record.Result = result.Output
				_ = SaveProcessingRecord(record)
				return APIResponse{RequestID: request.RequestID, Status: ResponseStatusSuccess, Output: result.Output}
			}
			continue
		}

		// Count failed REST calls and mark the agent unavailable at the limit.
		// 统计 REST 调用失败次数，并在达到上限时将智能体标记为不可用。
		record.FailureCount++
		if IncrementAgentFailureCount(agent.AgentID) >= agent.RetryLimit+1 {
			_ = UpdateAgentStatus(agent.AgentID, AgentStatusUnavailable)
			break
		}
	}

	// Save the final failure state and return an error response.
	// 保存最终失败状态并返回错误响应。
	record.Status = "failed"
	_ = SaveProcessingRecord(record)
	return APIResponse{RequestID: request.RequestID, Status: ResponseStatusError, ErrorMessage: "Agent request failed."}
}
```

### 4.2 Go Subfunctions
### 4.2 Go 子函数

Each function has an introduction comment for its purpose, input parameters, output parameters, and main flow.
每个函数都用介绍注释说明用途、输入参数、输出参数和主要流程。

```go
// ValidateRequest checks the request ID, agent ID, task text, and task length.
// ValidateRequest 检查请求 ID、智能体 ID、任务文本和任务长度。
// Input: request contains the caller's request fields.
// 输入：request 包含调用方提交的请求字段。
// Output: ValidationResult contains the valid flag and any error message.
// 输出：ValidationResult 包含有效标记和错误信息（如果有）。
func ValidateRequest(request Request) (result ValidationResult) {
	// Main flow: check required IDs, reject empty task text, and enforce the task length limit.
	// 主要流程：检查必需的 ID、拒绝空任务文本，并限制任务长度。
}

// GetSavedResponse finds a completed response for a repeated request ID.
// GetSavedResponse 根据重复请求 ID 查找已完成的响应。
// Input: requestID identifies the request.
// 输入：requestID 用于标识请求。
// Output: response, found flag, and any lookup error.
// 输出：响应、是否找到的标记，以及查询错误（如果有）。
func GetSavedResponse(requestID string) (response APIResponse, found bool, err error) {
	// Main flow: look up the existing processing result and map it to an API response.
	// 主要流程：查找已有的处理结果，并将其转换为 API 响应。
}

// TryAcquireRequestLock allows only one concurrent copy of a request to start processing.
// TryAcquireRequestLock 只允许一个并发请求副本开始处理。
// Input: requestID identifies the request lock.
// 输入：requestID 用于标识请求锁。
// Output: acquired flag and any lock error.
// 输出：是否成功获取锁的标记，以及锁错误（如果有）。
func TryAcquireRequestLock(requestID string) (acquired bool, err error) {
	// Main flow: create the request lock only if no copy already owns it.
	// 主要流程：仅当没有其他副本持有请求锁时创建该锁。
}

// WaitForRequestResult returns the response produced by the copy that owns the lock.
// WaitForRequestResult 返回持有请求锁的副本所生成的响应。
// Input: requestID identifies the in-progress request.
// 输入：requestID 用于标识正在处理的请求。
// Output: the completed API response.
// 输出：已完成的 API 响应。
func WaitForRequestResult(requestID string) (response APIResponse) {
	// Main flow: wait for completion and read the saved response.
	// 主要流程：等待处理完成并读取已保存的响应。
}

// ReleaseRequestLock releases the lock after request processing ends.
// ReleaseRequestLock 在请求处理结束后释放请求锁。
// Input: requestID identifies the lock to release.
// 输入：requestID 用于标识要释放的锁。
// Output: an error if the lock cannot be released.
// 输出：释放锁失败时返回错误。
func ReleaseRequestLock(requestID string) (err error) {
	// Main flow: remove the lock owned by this request processor.
	// 主要流程：移除由当前请求处理流程持有的锁。
}

// LoadAgent loads an agent's model settings and status from MySQL.
// LoadAgent 从 MySQL 加载智能体的模型设置和状态。
// Input: agentID identifies the requested agent.
// 输入：agentID 用于标识请求的智能体。
// Output: agent, found flag, and any lookup error.
// 输出：智能体对象、是否找到的标记，以及查询错误（如果有）。
func LoadAgent(agentID string) (agent Agent, found bool, err error) {
	// Main flow: read the agent record by ID.
	// 主要流程：根据 ID 读取智能体记录。
}

// SaveRequestRecord stores an accepted request in Redis.
// SaveRequestRecord 将已接收的请求保存到 Redis。
// Input: request contains the validated request data.
// 输入：request 包含已验证的请求数据。
// Output: an error if the record cannot be saved.
// 输出：记录保存失败时返回错误。
func SaveRequestRecord(request Request) (err error) {
	// Main flow: store the request fields under the request ID key.
	// 主要流程：使用请求 ID 作为键保存请求字段。
}

// SaveProcessingRecord stores or updates processing status, counts, and result in Redis.
// SaveProcessingRecord 在 Redis 中保存或更新处理状态、计数和结果。
// Input: record contains the current processing state.
// 输入：record 包含当前处理状态。
// Output: an error if the record cannot be saved.
// 输出：记录保存失败时返回错误。
func SaveProcessingRecord(record ProcessingRecord) (err error) {
	// Main flow: write the processing fields under the request ID key.
	// 主要流程：使用请求 ID 作为键写入处理字段。
}

// CallAgent sends task input to the agent REST API.
// CallAgent 将任务输入发送到智能体 REST API。
// Input: input contains the agent and task; timeoutSeconds sets the call timeout.
// 输入：input 包含智能体和任务；timeoutSeconds 设置调用超时时间。
// Output: AgentCallResult contains REST status, output, and any error message.
// 输出：AgentCallResult 包含 REST 状态、输出和错误信息（如果有）。
func CallAgent(input AgentCallInput, timeoutSeconds int) (result AgentCallResult) {
	// Main flow: prepare the API payload, call the agent, and return its response.
	// 主要流程：准备 API 请求内容、调用智能体并返回响应。
}

// ValidateAgentOutput checks whether the returned chat output is valid.
// ValidateAgentOutput 检查返回的聊天输出是否有效。
// Input: output is the agent's response text.
// 输入：output 是智能体返回的文本。
// Output: ValidationResult contains the valid flag and any error message.
// 输出：ValidationResult 包含有效标记和错误信息（如果有）。
func ValidateAgentOutput(output string) (result ValidationResult) {
	// Main flow: parse the output and check required response fields.
	// 主要流程：解析输出并检查必需的响应字段。
}

// IncrementAgentFailureCount atomically increments an agent's in-memory failure count.
// IncrementAgentFailureCount 以原子方式增加智能体的内存失败计数。
// Input: agentID identifies the agent's counter.
// 输入：agentID 用于标识智能体的计数器。
// Output: the updated consecutive failure count.
// 输出：更新后的连续失败次数。
func IncrementAgentFailureCount(agentID string) (consecutiveFailureCount int) {
	// Main flow: update the counter in the thread-safe map and return its new value.
	// 主要流程：更新线程安全映射中的计数器并返回新值。
}

// ClearAgentFailureCount removes an agent's in-memory failure key after a successful REST call.
// ClearAgentFailureCount 在 REST 调用成功后删除智能体的内存失败计数键。
// Input: agentID identifies the counter to remove.
// 输入：agentID 用于标识要删除的计数器。
// Output: no return value.
// 输出：无返回值。
func ClearAgentFailureCount(agentID string) {
	// Main flow: delete the agent ID key from the thread-safe map.
	// 主要流程：从线程安全映射中删除智能体 ID 键。
}

// UpdateAgentStatus writes the new agent status to MySQL.
// UpdateAgentStatus 将新的智能体状态写入 MySQL。
// Input: agentID identifies the agent; status is the new AgentStatus value.
// 输入：agentID 用于标识智能体；status 是新的 AgentStatus 值。
// Output: an error if the update fails.
// 输出：更新失败时返回错误。
func UpdateAgentStatus(agentID string, status AgentStatus) (err error) {
	// Main flow: update the status field for the selected agent.
	// 主要流程：更新指定智能体的状态字段。
}
```

## 5. Test Cases
## 5. 测试用例

Each case function returns sample input and expected output for one scenario.
每个用例函数都为一种场景返回样例输入和预期输出。

```go
type TestInput struct {
	Request              Request
	AgentExists          bool
	AgentStatus          AgentStatus
	RetryLimit           int
	ModelResults         []AgentCallResult
	InitialFailureCount  int
	ConcurrentCopies     int
	TaskTextLength       int
	TaskLengthLimit      int
}

type TestExpected struct {
	Response             APIResponse
	ProcessingStatus     string
	AttemptCount         int
	FailureCount         int
	AgentStatus          AgentStatus
	ModelCallCount       int
	FailureMapHasAgent   bool
	ResponsesMatch       bool
	RequestRecordSaved   bool
	ProcessingRecordSaved bool
}

type TestCase struct {
	Name     string
	Input    TestInput
	Expected TestExpected
}
```

### 5.1 Valid Request Succeeds
### 5.1 有效请求处理成功

```go
// CaseValidRequestSucceeds checks that a valid task reaches an available agent and returns its valid output.
// CaseValidRequestSucceeds 检查有效任务是否发送给可用智能体，并返回有效输出。
func CaseValidRequestSucceeds() TestCase {
	return TestCase{
		Name: "valid request succeeds",
		Input: TestInput{
			Request: Request{RequestID: "req-001", AgentID: "agent-001", TaskText: "Summarize this note."},
			AgentExists: true, AgentStatus: AgentStatusAvailable,
			ModelResults: []AgentCallResult{{RESTCallSucceeded: true, Output: "The note says the meeting is on Friday."}},
		},
		Expected: TestExpected{
			Response: APIResponse{RequestID: "req-001", Status: ResponseStatusSuccess, Output: "The note says the meeting is on Friday."},
			ProcessingStatus: "succeeded", AttemptCount: 1, ModelCallCount: 1,
		},
	}
}
```

### 5.2 Unknown Agent ID
### 5.2 智能体 ID 不存在

```go
// CaseUnknownAgentID checks that an unknown agent is rejected without a model call.
// CaseUnknownAgentID 检查系统是否拒绝未知智能体，且不调用模型。
func CaseUnknownAgentID() TestCase {
	return TestCase{
		Name: "unknown agent ID",
		Input: TestInput{
			Request: Request{RequestID: "req-002", AgentID: "agent-missing", TaskText: "Summarize this note."},
			AgentExists: false,
		},
		Expected: TestExpected{
			Response: APIResponse{RequestID: "req-002", Status: ResponseStatusFailure, ErrorMessage: "Agent does not exist."},
			ModelCallCount: 0,
		},
	}
}
```

### 5.3 Agent Is Unavailable
### 5.3 智能体不可用

```go
// CaseUnavailableAgent checks that an unavailable agent is rejected without a model call.
// CaseUnavailableAgent 检查系统是否拒绝不可用的智能体，且不调用模型。
func CaseUnavailableAgent() TestCase {
	return TestCase{
		Name: "agent unavailable",
		Input: TestInput{
			Request: Request{RequestID: "req-003", AgentID: "agent-002", TaskText: "Summarize this note."},
			AgentExists: true, AgentStatus: AgentStatusUnavailable,
		},
		Expected: TestExpected{
			Response: APIResponse{RequestID: "req-003", Status: ResponseStatusFailure, ErrorMessage: "Agent is unavailable."},
			ModelCallCount: 0,
		},
	}
}
```

### 5.4 Empty Task Text
### 5.4 任务文本为空

```go
// CaseEmptyTask checks that blank task text is rejected before calling a model.
// CaseEmptyTask 检查空白任务文本是否在调用模型前被拒绝。
func CaseEmptyTask() TestCase {
	return TestCase{
		Name: "empty task",
		Input: TestInput{
			Request: Request{RequestID: "req-004", AgentID: "agent-001", TaskText: "   "},
			AgentExists: true, AgentStatus: AgentStatusAvailable,
		},
		Expected: TestExpected{
			Response: APIResponse{RequestID: "req-004", Status: ResponseStatusFailure, ErrorMessage: "Task text is required."},
			ModelCallCount: 0,
		},
	}
}
```

### 5.5 Task Exceeds the Length Limit
### 5.5 任务超过长度限制

```go
// CaseTaskTooLong checks both the configured length boundary and rejection without a model call.
// CaseTaskTooLong 检查配置的长度边界，以及超长任务是否在调用模型前被拒绝。
func CaseTaskTooLong() TestCase {
	return TestCase{
		Name: "task exceeds length limit",
		Input: TestInput{
			Request: Request{RequestID: "req-005", AgentID: "agent-001", TaskText: "<task text of 10,001 characters>"},
			AgentExists: true, AgentStatus: AgentStatusAvailable,
			TaskTextLength: 10001, TaskLengthLimit: 10000,
		},
		Expected: TestExpected{
			Response: APIResponse{RequestID: "req-005", Status: ResponseStatusFailure, ErrorMessage: "Task text exceeds the length limit."},
			ModelCallCount: 0,
		},
	}
}
```

### 5.6 Concurrent Copies of the Same Request
### 5.6 同一请求的并发副本

```go
// CaseConcurrentDuplicateRequest checks that concurrent copies share one model call and one result.
// CaseConcurrentDuplicateRequest 检查并发副本是否共用一次模型调用和同一个结果。
func CaseConcurrentDuplicateRequest() TestCase {
	return TestCase{
		Name: "concurrent duplicate request",
		Input: TestInput{
			Request: Request{RequestID: "req-006", AgentID: "agent-001", TaskText: "Summarize this note."},
			AgentExists: true, AgentStatus: AgentStatusAvailable, ConcurrentCopies: 2,
			ModelResults: []AgentCallResult{{RESTCallSucceeded: true, Output: "The note says the meeting is on Friday."}},
		},
		Expected: TestExpected{
			Response: APIResponse{RequestID: "req-006", Status: ResponseStatusSuccess, Output: "The note says the meeting is on Friday."},
			ProcessingStatus: "succeeded", AttemptCount: 1, ModelCallCount: 1, ResponsesMatch: true,
		},
	}
}
```

**Test note:** Submit both copies concurrently with the same request ID; check the atomic Redis lock and saved result.
**测试说明：** 使用相同请求 ID 并发提交两个副本；检查 Redis 原子锁和已保存结果。

### 5.7 Model Call Fails Until the Retry Limit
### 5.7 模型调用失败直至重试上限

```go
// CaseModelCallRetryLimit checks retry count, failure recording, and the agent status change after repeated timeouts.
// CaseModelCallRetryLimit 检查多次超时后的重试次数、失败记录和智能体状态变更。
func CaseModelCallRetryLimit() TestCase {
	timeout := AgentCallResult{RESTCallSucceeded: false, ErrorMessage: "model call timed out"}
	return TestCase{
		Name: "model call reaches retry limit",
		Input: TestInput{
			Request: Request{RequestID: "req-007", AgentID: "agent-001", TaskText: "Summarize this note."},
			AgentExists: true, AgentStatus: AgentStatusAvailable, RetryLimit: 2,
			ModelResults: []AgentCallResult{timeout, timeout, timeout},
		},
		Expected: TestExpected{
			Response: APIResponse{RequestID: "req-007", Status: ResponseStatusError, ErrorMessage: "Agent request failed."},
			ProcessingStatus: "failed", AttemptCount: 3, FailureCount: 3,
			AgentStatus: AgentStatusUnavailable, ModelCallCount: 3, FailureMapHasAgent: true,
		},
	}
}
```

**Test note:** A retry limit of two allows three total calls. Also check thread-safe counter updates and counter deletion after a successful REST call.
**测试说明：** 重试上限为 2 时总共允许调用 3 次。还要检查计数器的并发安全，以及 REST 调用成功后是否删除计数。

### 5.8 Successful REST Call Clears the Failure Counter
### 5.8 REST 调用成功后清除失败计数

```go
// CaseSuccessClearsFailureCount checks that a successful REST call deletes the agent's in-memory failure key.
// CaseSuccessClearsFailureCount 检查 REST 调用成功后是否删除智能体的内存失败计数键。
func CaseSuccessClearsFailureCount() TestCase {
	return TestCase{
		Name: "success clears failure count",
		Input: TestInput{
			Request: Request{RequestID: "req-008", AgentID: "agent-001", TaskText: "Summarize this note."},
			AgentExists: true, AgentStatus: AgentStatusAvailable, InitialFailureCount: 2,
			ModelResults: []AgentCallResult{{RESTCallSucceeded: true, Output: "The note says the meeting is on Friday."}},
		},
		Expected: TestExpected{
			Response: APIResponse{RequestID: "req-008", Status: ResponseStatusSuccess, Output: "The note says the meeting is on Friday."},
			ProcessingStatus: "succeeded", AttemptCount: 1, ModelCallCount: 1, FailureMapHasAgent: false,
		},
	}
}
```

### 5.9 Invalid Model Output Reaches the Retry Limit
### 5.9 模型输出无效并达到重试上限

```go
// CaseInvalidOutputRetryLimit checks output validation, retry behavior, failure details, and final agent availability.
// CaseInvalidOutputRetryLimit 检查输出校验、重试行为、失败详情和最终智能体可用状态。
func CaseInvalidOutputRetryLimit() TestCase {
	invalidOutput := AgentCallResult{RESTCallSucceeded: true, Output: ""}
	return TestCase{
		Name: "invalid output reaches retry limit",
		Input: TestInput{
			Request: Request{RequestID: "req-009", AgentID: "agent-001", TaskText: "Summarize this note."},
			AgentExists: true, AgentStatus: AgentStatusAvailable, RetryLimit: 1,
			ModelResults: []AgentCallResult{invalidOutput, invalidOutput},
		},
		Expected: TestExpected{
			Response: APIResponse{RequestID: "req-009", Status: ResponseStatusError, ErrorMessage: "Agent request failed."},
			ProcessingStatus: "failed", AttemptCount: 2, AgentStatus: AgentStatusUnavailable, ModelCallCount: 2,
		},
	}
}
```

**Test note:** REST succeeds in this case, but its content is invalid. Verify that the workflow still retries, records the validation failure, and applies the documented final agent status.
**测试说明：** 此用例中的 REST 调用成功，但返回内容无效。验证流程仍会重试、记录校验失败，并按设计更新最终智能体状态。

### 5.10 Stored Request and Processing Records
### 5.10 已保存的请求和处理记录

```go
// CaseRecordsPersisted checks that a successful request and its processing result are saved in Redis.
// CaseRecordsPersisted 检查成功请求及其处理结果是否保存到 Redis。
func CaseRecordsPersisted() TestCase {
	return TestCase{
		Name: "request and processing records are saved",
		Input: TestInput{
			Request: Request{RequestID: "req-010", AgentID: "agent-001", TaskText: "Summarize this note."},
			AgentExists: true, AgentStatus: AgentStatusAvailable,
			ModelResults: []AgentCallResult{{RESTCallSucceeded: true, Output: "The note says the meeting is on Friday."}},
		},
		Expected: TestExpected{
			Response: APIResponse{RequestID: "req-010", Status: ResponseStatusSuccess, Output: "The note says the meeting is on Friday."},
			ProcessingStatus: "succeeded", AttemptCount: 1, ModelCallCount: 1,
			RequestRecordSaved: true, ProcessingRecordSaved: true,
		},
	}
}
```

**Test note:** Confirm record fields use the matching request ID and do not include secrets or unrelated sensitive data.
**测试说明：** 确认记录字段使用对应的请求 ID，且不包含密码等无关敏感数据。
