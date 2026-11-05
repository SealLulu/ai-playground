# 0. [Project Name] System Design
# 0. [项目名称]系统设计

## 0.0 Rules of This File
## 0.0 本文件的规则

Use English first and Mandarin second for all content.
所有内容先使用英语，后使用中文。

Use simple and clear sentences.
使用简单、清晰的句子。

Start every heading with its index number.
每个标题以其索引编号开头。

Replace bracketed placeholders with project-specific content.
将方括号占位符替换为项目特定内容。

Keep all sections consistent and update them when the design changes.
保持所有部分一致，并在设计发生变化时更新它们。

Keep required topics and select optional topics according to the current project.
保留必选主题，并根据当前项目选择可选主题。

## 1. Project Summary
## 1. 项目概述

## 1.0 Rule
## 1.0 要求

Before writing any subsection in Section 1, ask the user what content they want in that subsection, provide your own suggested additions based on the project requirements, and wait for the user's confirmation. Write the subsection only after the user confirms its content.

在编写第 1 节中的任何子节之前，先询问用户希望该子节包含哪些内容，根据项目需求提出你建议补充的内容，并等待用户确认。只有在用户确认内容后，才能编写该子节。

### 1.1 Business Summary
### 1.1 业务概述

### 1.2 Functional Requirements
### 1.2 功能需求

### 1.2.0 Rule
### 1.2.0 要求

Identify the project type and its core design requirements first.

首先识别项目类型和核心设计要求。

### 1.3 Non-Functional Requirements
### 1.3 非功能需求

### 1.4 Difficult Problems and Edge Cases
### 1.4 难点和边界情况

#### 1.4.0 Rule
#### 1.4.0 要求

Every project must cover concrete problems related to validation, core entity and state lifecycle, and failure handling.

每个项目都必须覆盖与验证、核心实体和状态生命周期以及故障处理相关的具体问题。

Also consider duplicate and idempotent operations, concurrent operations, resource availability, lock order and deadlock prevention, recovery, context management, streaming, and external service integration when they are relevant to the current project.

当重复和幂等操作、并发操作、资源可用性、锁顺序和死锁预防、恢复、上下文管理、流式处理以及外部服务集成与当前项目相关时，也要考虑这些方面。

Use these categories only to discover and verify coverage. Do not use category or definition names such as `Validation`, `Core Entity and State Lifecycle`, `Failure Handling`, or `Concurrent Operations` as final subsection headings.

这些类别仅用于发现问题和检查覆盖范围。不要将“验证”、“核心实体和状态生命周期”、“故障处理”或“并发操作”等类别或定义名称用作最终子标题。

Give every subsection a sequential bilingual heading that names one actual project problem, conflict, failure event, or edge case, such as `Duplicate Requests`, `Concurrent Requests`, or `Server Exit During Processing`.

为每个子节使用连续编号的双语标题，并在标题中写明一个实际项目问题、冲突、故障事件或边界情况，例如“重复请求”、“并发请求”或“服务器在处理期间退出”。

Keep each heading specific enough to identify what happened or what can go wrong. Describe the handling rule, expected behavior, and correctness guarantee under that heading.

每个标题都必须足够具体，能够说明发生了什么或可能出现什么问题。在该标题下描述处理规则、预期行为和正确性保证。

#### 1.4.1 [First Actual Problem or Edge Case]
#### 1.4.1 [第一个实际问题或边界情况]

## 2. Business Logic
## 2. 业务逻辑

### 2.1 Components
### 2.1 组件

#### 2.1.0 Rule
#### 2.1.0 要求

Add one indexed component heading for each component required by the current project.

为当前项目需要的每个组件添加一个带索引的组件标题。

Under each component, describe every relevant function as an indexed workflow. Use a Mermaid `flowchart TD` diagram for each workflow.

在每个组件下，将每个相关功能描述为带索引的工作流。每个工作流都使用 Mermaid `flowchart TD` 流程图。

Use English first and Mandarin second in workflow headings, nodes, branch labels, and explanations. In Mermaid nodes and branch labels, place Mandarin after English with `<br/>`.

工作流标题、节点、分支标签和解释先使用英语，后使用中文。在 Mermaid 节点和分支标签中，使用 `<br/>` 将中文放在英文之后。

After a workflow diagram, add concise bilingual explanations only when they clarify a guarantee, invariant, validation rule, concurrency rule, or prevented error.

仅当需要说明保证条件、不变量、验证规则、并发规则或所防止的错误时，才在工作流图后添加简洁的双语解释。

Do not create a separate Component Responsibilities section. Express responsibilities through each component's functional workflows and their explanations.

不要创建单独的“组件职责”章节。通过每个组件的功能工作流及其解释来表达组件职责。

Describe Section 2.2 as an end-to-end Mermaid `flowchart TD` workflow. Use English first and Mandarin second in every node and branch label, separated with `<br/>`.

将第 2.2 节描述为端到端的 Mermaid `flowchart TD` 工作流。每个节点和分支标签先使用英语，后使用中文，并用 `<br/>` 分隔。

#### 2.1.1 [First Component]
#### 2.1.1 [第一个组件]

##### 2.1.1.1 [First Function Workflow]
##### 2.1.1.1 [第一个功能工作流]

### 2.2 Main Processing Workflow
### 2.2 主要处理流程

## 3. Storage
## 3. 存储

### 3.1 Main Storage Structure
### 3.1 主要存储结构

#### 3.1.0 Rule
#### 3.1.0 要求

Before writing Section 3.1, identify the foundational storage, data-processing, and synchronization choices required by the main process in Section 2.2. Discuss the recommended choices with the user and wait for confirmation before writing this section.

在编写第 3.1 节之前，识别第 2.2 节主要处理流程需要的基础存储、数据处理和同步选型。与用户讨论建议的选型，并在编写本节之前等待用户确认。

Select only mechanisms required by the current project. Examples include a hash map, min-heap, queue, streaming or channel mechanism, mutex, read-write mutex, and atomic operation.

只选择当前项目需要的机制。例如哈希表、最小堆、队列、流式处理或通道机制、互斥锁、读写锁以及原子操作。

Use sequential headings named after the basic technical choice itself, such as `Hash Map`, `Min-Heap`, `Streaming`, `Mutex and RWMutex`, or `Atomic Operations`.

使用以基础技术选型本身命名的连续编号标题，例如“哈希表”、“最小堆”、“流式处理”、“互斥锁和读写锁”或“原子操作”。

Do not use a business problem, failure event, or edge case as a Section 3.1 heading. Problems such as duplicate requests, concurrent requests, or server exit belong in Sections 1.4 and 2.

不要将业务问题、故障事件或边界情况用作第 3.1 节标题。重复请求、并发请求或服务器退出等问题属于第 1.4 节和第 2 节。

For each choice, explain where the main process uses it, what data or state it manages, why it is selected, its key operations or ordering rules, and any required update invariant.

对于每个选型，说明主要处理流程在何处使用它、它管理哪些数据或状态、选择它的原因、关键操作或排序规则以及必须保持的更新不变量。

For a lock or atomic mechanism, also explain the protected state, lock mode or atomic boundary, acquisition and release rules, and interaction with other selected mechanisms.

对于锁或原子机制，还要说明受保护的状态、锁模式或原子边界、获取和释放规则，以及它与其他已选机制的交互方式。

Do not repeat business-unit field definitions from Section 3.2, and do not introduce unused storage or synchronization choices.

不要重复第 3.2 节中的业务单元字段定义，也不要引入未使用的存储或同步选型。

#### 3.1.1 [First Basic Storage, Processing, or Synchronization Choice]
#### 3.1.1 [第一个基础存储、处理或同步选型]

### 3.2 Business Unit Storage Structures
### 3.2 业务单元存储结构

#### 3.2.0 Rule
#### 3.2.0 要求

Add storage structures only for business units relevant to the current project.

只为与当前项目相关的业务单元添加存储结构。

Use one business-specific heading for each business unit structure. Do not output placeholder or generic structure headings.

为每个业务单元结构使用一个业务特定标题。不要输出占位符或通用结构标题。

Define every business unit structure in a fenced Go code block.

在带围栏的 Go 代码块中定义每个业务单元结构。

Before each related type, define the constants required by its business states, fixed limits, default values, or other reused fixed values. Do not create constants that the business logic does not need.

在每个相关类型之前，定义其业务状态、固定限制、默认值或其他重复使用的固定值所需的常量。不要创建业务逻辑不需要的常量。

Add a concise end-of-line Mandarin comment after every field that carries business meaning, state, counters, indexes, flags, limits, timestamps, or relationships. The comment must explain the field's business purpose rather than repeat its name.

在每个承载业务含义、状态、计数器、索引、标志、限制、时间或关联关系的字段后添加简洁的中文行尾注释。注释必须说明字段的业务用途，而不是重复字段名称。

Keep the Go definitions minimal, syntactically consistent, and aligned with the business logic and storage design. Do not add methods or implementation logic in Section 3.2.

保持 Go 定义简洁、语法一致，并与业务逻辑和存储设计保持一致。不要在第 3.2 节中添加方法或实现逻辑。

#### 3.2.1 [First Business Unit Structure]
#### 3.2.1 [第一个业务单元结构]

## 4. Coding
## 4. 编码

### 4.0 Rule
### 4.0 要求

Use Go and implement only the core functions required by the current task.

使用 Go，并且只实现当前任务所需的核心函数。

Keep supporting functions minimal.

保持辅助函数简洁。

Define the required data structures and function signatures before the main function.

在主函数之前定义所需的数据结构和函数签名。

The main function must show only the main flow and call minimal supporting functions.

主函数必须只展示主要流程，并调用最少的辅助函数。

Use TODO comments instead of detailed unfinished logic.

使用 TODO 注释代替详细的未完成逻辑。

Describe each supporting function's purpose, inputs, and return value, and keep all signatures and calls consistent.

说明每个辅助函数的用途、输入和返回值，并保持所有函数签名和调用一致。

For every function or method, place a bilingual comment block immediately above the `func` declaration. State the function's purpose, explain every input parameter, and explain every output value, including errors. State explicitly when there is no input or output.

对于每个函数或方法，在 `func` 声明正上方放置双语注释块。说明函数用途，解释每个输入参数，并解释每个输出值，包括错误。没有输入或输出时也要明确说明。

Write the function summary first, followed by `Input:` and `Output:` comments. Put each English comment before its corresponding Mandarin comment.

先编写函数摘要，然后编写 `Input:` 和 `Output:` 注释。每条英文注释放在对应的中文注释之前。

Inside every function, place concise bilingual comments immediately before each important validation, state transition, storage update, lock operation, external call, error path, or other key processing step. Explain the purpose and business effect of the step, not merely the syntax.

在每个函数内部，在每个重要的验证、状态转换、存储更新、锁操作、外部调用、错误路径或其他关键处理步骤之前放置简洁的双语注释。说明该步骤的目的和业务效果，而不是仅仅描述语法。

Add a concise bilingual end-of-line comment to every data-structure field that carries business meaning. Put English before Mandarin in the same comment.

为每个承载业务含义的数据结构字段添加简洁的双语行尾注释。在同一条注释中先写英文，后写中文。

When concurrency rules apply, write each applicable locking or atomicity rule as paired English and Mandarin comments before the affected function or code block. Include lock-order, incompatible-lock, release-before-blocking-operation, and atomic-boundary rules only when required by the confirmed design.

当并发规则适用时，在受影响的函数或代码块之前，将每条适用的锁定或原子性规则写成英文和中文配对注释。仅在已确认的设计需要时，包含锁顺序、不兼容锁、阻塞操作前释放锁以及原子边界规则。

Do not invent comments or concurrency rules that are not supported by the design. Keep comments specific, useful, and consistent with the code.

不要编造设计中不存在的注释或并发规则。保持注释具体、有用，并与代码一致。

### 4.1 Main Process Code
### 4.1 主流程代码

#### 4.1.1 Main Function
#### 4.1.1 主函数

#### 4.1.2 Sub-Function Definitions
#### 4.1.2 子函数定义

### 4.2 Business Unit Code
### 4.2 业务单元代码

#### 4.2.0 Rule
#### 4.2.0 要求

Add code only for the business units required by the current project.

只添加当前项目需要的业务单元代码。

Each business unit must show only its main processing flow.

每个业务单元必须只展示其主要处理流程。

Do not add detailed implementation logic to the business unit process.

不要在业务单元流程中添加详细的实现逻辑。

For each key step, add paired English and Mandarin comments first and then call its sub-function.

对于每个关键步骤，先添加英文和中文配对注释，然后调用其子函数。

#### 4.2.1 [First Business Unit Code]
#### 4.2.1 [第一个业务单元代码]

### 4.3 Error and Concurrency Handling
### 4.3 错误和并发处理

## 5. Test Data
## 5. 测试数据

### 5.0 Rule
### 5.0 要求

Cover functional requirements, non-functional requirements, validation, core entity and state lifecycle, and failure handling in every project. Add risk-specific coverage only when it is relevant to the confirmed design.

每个项目都必须覆盖功能需求、非功能需求、验证、核心实体和状态生命周期以及故障处理。仅在与已确认设计相关时添加特定风险覆盖。

Use these coverage categories only to select tests. Do not use generic category names as test headings.

这些覆盖类别仅用于选择测试。不要将通用类别名称用作测试标题。

Give every test scenario a sequential heading that names the actual problem or behavior under test, such as `Invalid Input` or `Concurrent Requests`.

为每个测试场景使用连续编号标题，并在标题中写明实际测试的问题或行为，例如“非法输入”或“并发请求”。

Every test scenario must contain `Input` and `Expected Response` subsections. Encode both subsections as Go test data in fenced `go` code blocks. Use composite literals, slices, maps, constants, variables, or error values that match the structures and identifiers defined in Sections 3 and 4.

每个测试场景都必须包含“输入”和“预期响应”子节。两个子节都使用带围栏的 `go` 代码块编码为 Go 测试数据。使用与第 3 节和第 4 节中定义的结构及标识符一致的复合字面量、切片、映射、常量、变量或错误值。

Add concise bilingual Go comments to important fields and values. Put English first and Mandarin second, and explain the field's test purpose or significance rather than repeating its name.

为重要字段和值添加简洁的双语 Go 注释。先写英语，后写中文，并说明字段的测试用途或重要性，而不是重复字段名称。

Within each Go struct literal or ordered test-data definition, list all required business fields first. If optional fields are included, append them after the required fields in their business order.

在每个 Go 结构体字面量或有序测试数据定义中，先列出所有必选业务字段。如果包含可选字段，则按照业务顺序将其接在必选字段之后。

Do not add optional fields that are unrelated to the test scenario. Keep Input and Expected Response consistent with the requirements, workflows, storage, and code.

不要添加与测试场景无关的可选字段。保持输入和预期响应与需求、工作流、存储和代码一致。

### 5.1 [First Actual Test Problem]
### 5.1 [第一个实际测试问题]

#### 5.1.1 Input
#### 5.1.1 输入

#### 5.1.2 Expected Response
#### 5.1.2 预期响应
