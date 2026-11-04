# Multi-Agent
# Multi-Agent 多智能体项目

Multi-Agent routes a request to an available agent, calls its model, validates the response, and records the processing result.
Multi-Agent 将请求路由给可用的智能体，调用其模型，校验响应，并记录处理结果。

This repository contains a runnable Go prototype based on the system design in `Multi-Agent.md`.
本仓库包含一个可运行的 Go 原型，设计依据见 `Multi-Agent.md`。

## Project Files
## 项目文件

| File | Purpose |
|---|---|
| [agent.go](agent.go) | Defines agents, availability status, the agent repository interface, a thread-safe in-memory repository, and per-agent failure counters. / 定义智能体、可用状态、智能体仓库接口、线程安全的内存仓库，以及每个智能体的失败计数。 |
| [request.go](request.go) | Defines API request and response structures, status codes, and request validation for required fields and task length. / 定义 API 请求和响应结构、状态码，以及必填字段和任务长度校验。 |
| [model.go](model.go) | Defines the model-call interface, timeout handling, and a local echo model for the runnable demo. / 定义模型调用接口、超时处理，以及供演示程序使用的本地回显模型。 |
| [record.go](record.go) | Defines request and processing records, the record-store interface, and a thread-safe in-memory store with duplicate-request coordination. / 定义请求和处理记录、记录存储接口，以及支持并发重复请求协调的线程安全内存存储。 |
| [processing.go](processing.go) | Implements request validation, agent lookup, duplicate-request handling, model retries, output validation, failure tracking, and response creation. / 实现请求校验、智能体查找、重复请求处理、模型重试、输出校验、失败计数和响应生成。 |
| [processing_test.go](processing_test.go) | Tests the request-processing scenarios described in Section 5 of the design. / 测试设计文档第 5 节中的请求处理场景。 |
| [cmd/multiagent/main.go](cmd/multiagent/main.go) | Starts the local demo, submits a sample task, and prints the API response as JSON. / 启动本地演示，提交示例任务，并以 JSON 格式打印 API 响应。 |
| [go.mod](go.mod) | Declares the Go module name and minimum Go version. / 声明 Go 模块名称和最低 Go 版本。 |
| [Multi-Agent.md](Multi-Agent.md) | Describes the bilingual system design, business logic, storage model, processing outline, and test scenarios. / 描述双语系统设计、业务逻辑、存储模型、处理流程框架和测试场景。 |
| [Agent-Call-Workflow.drawio](Agent-Call-Workflow.drawio) | Editable diagram for the agent call and response-validation workflow. / 智能体调用和响应校验流程的可编辑图表。 |
| [Main-Processing-Workflow.drawio](Main-Processing-Workflow.drawio) | Editable diagram for the main request-processing workflow. / 主请求处理流程的可编辑图表。 |
| [TEST-REPORT.md](TEST-REPORT.md) | Records the latest test results and commands used to run the demo. / 记录最近的测试结果和运行演示程序的命令。 |

## Run
## 运行

Run all tests:
运行全部测试：

```powershell
go test ./...
```

Run the demo:
运行演示程序：

```powershell
go run ./cmd/multiagent
```

The demo uses in-memory agent and record stores and a local echo model, so it does not require MySQL, Redis, or external model credentials.
演示程序使用内存版智能体和记录存储，以及本地回显模型，因此不需要 MySQL、Redis 或外部模型凭据。

## Storage and Model Adapters
## 存储和模型适配器

The repository and model interfaces are integration points for MySQL, Redis, and a real model REST API. The current memory implementations support local runs and tests; persistent database and Redis adapters are not included yet.
仓库和模型接口用于接入 MySQL、Redis 和真实模型 REST API。目前的内存实现用于本地运行和测试；项目尚未包含持久化数据库和 Redis 适配器。
