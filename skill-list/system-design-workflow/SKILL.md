---
name: system-design-workflow
description: Create new bilingual system design documents using the bundled template. Construct documents progressively, generating headings and content only when requested and confirmed. Always follow the template's global rules without outputting them. Use only for new system designs, not modifications to existing projects or established designs.
---

# System Design Workflow

Use the bundled `assets/SYSTEM_DESIGN_TEMPLATE.md` as the canonical specification for system design structure, global rules, requirements, and authoring constraints.

## 1. Global Template Rules

### 1.1 Global Rules in Section 0.0

`0.0 Rules of This File` defines global rules for the entire System Design Document.

These rules are mandatory and apply to ALL generated content.

The Agent MUST:

1. Read and internalize Section 0.0 before creating or modifying a project design.
2. Follow every applicable rule in Section 0.0 throughout the entire workflow.
3. Never generate Section 0.0 in the project design document.
4. Never copy the English or Mandarin instructions under Section 0.0 into the project design document.
5. Apply the global rules to every generated heading, paragraph, code section, and test case.
6. Preserve the project title format: `# 0. [Project Name] System Design`.
7. Use the section numbering defined by the current template.

The global rules are not optional, even if the user requests only one subsection.

### 1.2 Global Content Requirements

All generated project content MUST follow these requirements:

- English first, Mandarin second.
- Simple and clear sentences.
- Every heading starts with its correct index number.
- Replace bracketed placeholders with project-specific content.
- Keep all sections consistent when the design changes.
- Preserve required topics.
- Select optional topics according to the current project's requirements and risks.

These requirements apply to all sections, including business logic, storage, coding, and test data.

## 2. Trigger Conditions

Use this skill when the user:

- Requests a new system design file for a new system.
- Asks to design a system from scratch.
- Asks to start or build a new system design using the bundled template.
- Asks to fill or complete sections as part of the same new-system-design workflow.
- Explicitly invokes `$system-design-workflow` or `/system-design-workflow` to start a new system design.

Do not use this skill when the user:

- Requests a code change, bug fix, or feature change in an existing project.
- Requests an architecture or behavior change in an existing system.
- Requests an update to an established system design document.
- Requests only code implementation or a general architecture explanation.
- Requests an unrelated document or template.

## 3. Workflow Boundary

### 3.1 New System Design

Use this skill to create a new design file from the bundled template and complete it progressively, section by section.

### 3.2 Existing Project Modification

Do not use this skill.

Follow the existing project's code, architecture, requirements, and user instructions without forcing the system design template.

## 4. Template Rules

1. The bundled template is the canonical source of truth.
2. Always follow Section 0.0.
3. Follow the current template's section numbering and hierarchy.
4. Preserve all applicable required topics.
5. Select optional topics according to the project's requirements and risks.
6. Preserve explicit user constraints.
7. Do not invent requirements that are not specified by the user or reasonably required by the project.
8. Do not modify the bundled template when creating or updating a project design.
9. If the user explicitly requests a change to the template itself, treat it as a separate template-editing task.
10. Complete only the sections requested by the user during the normal section-by-section workflow.
11. Preserve existing user content outside the requested scope.

## 5. Progressive Document Construction

The final System Design Document MUST be constructed progressively.

Do NOT generate the complete template hierarchy at the beginning.

### 5.1 Initial Document Creation

When creating a new System Design Document:

1. Read the bundled template completely.
2. Read and internalize Section 0.0.
3. Create a new Markdown file.
4. Replace the project-name placeholder if the project name is known.
5. Generate ONLY the outermost project design headings:

   - `# 0. [Project Name] System Design`
   - `## 1. Project Summary`
   - `## 2. Business Logic`
   - `## 3. Storage`
   - `## 4. Coding`
   - `## 5. Test Data`

6. Do not generate any deeper headings.
7. Do not generate subsection content.
8. Do not generate required or optional topic lists.
9. Do not generate `0.0 Rules of This File`.
10. Do not copy any global rules or internal instructions into the project document.

The initial document must contain only the project title and the five outermost design sections.

### 5.2 Generate Headings Only When Needed

Apply progressive construction recursively to every section.

When the user requests a section:

1. Identify the requested section and its position in the template.
2. Read the corresponding template rules.
3. Identify the immediate child headings required for the requested work.
4. Generate only the necessary headings.
5. Do not generate unrelated sibling headings.
6. Do not generate deeper headings until their corresponding subsections are requested and applicable decisions are confirmed.
7. Add content only after satisfying the confirmation requirements of that section.

Example:

If the user requests Section 1.1:

- Generate `### 1.1 Business Summary`.
- Do not generate Sections 1.2, 1.3, or 1.4.
- Do not generate deeper headings.
- Follow the confirmation requirement in Section 1.0.
- Write the subsection only after user confirmation.

If the user later requests Section 1.4:

- Generate `### 1.4 Difficult Problems and Edge Cases`.
- Identify the required coverage categories and relevant risk categories.
- Derive concrete project problems and edge cases from those categories.
- Discuss the proposed concrete problem list with the user.
- Wait for confirmation.
- Generate only the confirmed actual-problem headings.
- Generate deeper headings only when needed.

Apply this rule recursively to all heading levels.

### 5.3 Required and Optional Topic Selection

Whenever a subsection contains required and optional topics, discuss the topic list with the user before generating the corresponding topic headings.

The process MUST be:

1. Read all required and optional topics from the template.
2. Present all required topics to the user.
3. Explain that required topics must be included.
4. Allow the user to discuss their scope, content, and wording.
5. Analyze the project and recommend relevant optional topics.
6. Explain briefly why each recommended optional topic is relevant.
7. Present the proposed final topic list.
8. Wait for the user's confirmation.
9. Generate the confirmed topic headings.
10. Add detailed content when the user requests the corresponding topics.

Rules:

- Required topics must always be preserved.
- Do not silently remove required topics.
- User confirmation determines the scope and wording, but does not silently waive mandatory topics.
- Optional topics must not be generated without confirmation.
- Do not generate every optional topic by default.
- Do not generate empty headings for unselected optional topics.
- Do not generate deeper headings for topics that have not been selected.
- If the user changes the selected list, update the corresponding headings and dependent sections.
- For Section 1.4, required and optional topics are coverage categories only. Convert them into concrete project problems before generating headings.
- Never output a Section 1.4 category name as a heading when a more specific actual problem can be named.

### Optional Topic Numbering

- Concrete problems derived from optional categories must continue the numbering sequence after problems that satisfy required coverage.
- Do not restart numbering for optional topics.
- Treat Required Topics and Optional Topics as internal grouping labels, not as separate numbering branches in the final document.
- Keep problems that satisfy required coverage first, followed by confirmed problems derived from optional risk categories.
- Renumber every confirmed actual-problem heading sequentially without category-based numbering branches.
- Do not leave numbering gaps for optional topics that were not selected.
- Apply this rule consistently to all sections containing required and optional topics.

### 5.4 Internal Rules and Organizational Headings

The following template headings are internal instructions or topic-selection containers. They MUST NOT appear in the final project design document:

- `0.0 Rules of This File`
- All `.0 Rule` headings and their instructional paragraphs.

Use internal rules and coverage categories to guide topic selection.

Generate only confirmed actual-problem or project-specific topic headings directly under the relevant document section.

For example, derive headings such as `Invalid Request Data`, `Duplicate Requests`, `Concurrent Request Registration`, or `Server Exit During Processing` from the internal coverage categories.

Do not generate the organizational headings Required Topics or Optional Topics.

### 5.5 No Premature Heading Generation

Never generate the full hierarchy of a section simply because its parent section has been created.

For example, creating Section 3 does not authorize generating:

- `3.1 Main Storage Structure`
- `3.2 Business Unit Storage Structures`
- Any basic storage or data-processing choice under Section 3.1
- Any synchronization or atomic-operation choice under Section 3.1

These headings must be generated progressively when their corresponding sections are requested and applicable decisions are confirmed.

## 6. Section-Specific Workflow

### 6.1 Section 1: Project Summary

Before writing any subsection in Section 1:

1. Identify the project type and its core design requirements.
2. Ask the user what content they want in the requested subsection.
3. Provide your own suggested additions based on the project requirements.
4. Wait for the user's confirmation.
5. Generate the requested subsection heading.
6. Generate necessary child headings only when required.
7. Write the subsection only after confirmation.

Do not skip the confirmation step even when the project requirements seem obvious.

For Section 1.2:

- Identify the project type first.
- Identify its core design requirements.
- Discuss the proposed functional requirements with the user.
- Generate only the confirmed content structure.

For Section 1.4:

- Treat Validation, Core Entity and State Lifecycle, and Failure Handling as required coverage categories, not final headings.
- Explain that concrete problems covering all three categories must be included.
- Consider Duplicate and Idempotent Operations, Concurrent Operations, Resource Availability, Lock Order and Deadlock Prevention, Recovery, Context Management, Streaming, and External Service Integration as risk categories when relevant.
- Analyze the project and derive concrete problems, conflicts, failure events, and edge cases from the applicable categories.
- Give each proposed item a specific bilingual problem name, such as `Duplicate Requests` / `重复请求`, `Concurrent Requests` / `并发请求`, or `Server Exit During Processing` / `服务器在处理期间退出`.
- Do not propose category or definition names such as `Validation`, `Failure Handling`, or `Concurrent Operations` as final headings.
- Present the proposed concrete problem list and explain which required or risk categories it covers.
- Wait for confirmation.
- Generate the confirmed problems as sequential `1.4.x` headings with English first and Mandarin second.
- Under each heading, describe the handling rule, expected behavior, and correctness guarantee for that specific problem.
- Do not generate organizational headings for required or optional categories.

### 6.2 Section 2: Business Logic

Section 2 contains only these project-document subsections:

- `2.1 Components`
- `2.2 Main Processing Workflow`

For Section 2.1:

- Treat `2.1.0 Rule` in the template as an internal authoring rule. Never output that heading or its instructions in the project document.
- Add one sequentially numbered component heading for each actual project component, such as `2.1.1`, `2.1.2`, and so on.
- Under each component, add one sequentially numbered workflow heading for every relevant component function, such as `2.1.1.1`, `2.1.1.2`, and so on.
- Describe every component function with a Mermaid `flowchart TD` diagram.
- Use English first and Mandarin second in workflow headings, diagram nodes, branch labels, and explanatory text.
- In Mermaid nodes and branch labels, separate English and Mandarin with `<br/>`.
- After a diagram, add concise bilingual explanations when they clarify a guarantee, invariant, validation rule, concurrency rule, or prevented error.
- Express component responsibilities through the functional workflows and explanations. Do not create a separate Component Responsibilities subsection.
- Generate only components and functional workflows that are relevant to the current project. Do not output placeholder components.

For Section 2.2:

- Describe the end-to-end main process as a Mermaid `flowchart TD` workflow.
- Use English first and Mandarin second in every node and branch label, separated with `<br/>`.
- Keep the main workflow consistent with the component workflows, requirements, and storage design.

Generate only the requested subsection. Do not generate both Section 2.1 and Section 2.2 unless the user requests the complete Section 2.

### 6.3 Section 3.1: Main Storage Structure

Before writing Section 3.1:

1. Read the main process in Section 2.2 and identify the foundational storage, data-processing, and synchronization mechanisms it actually needs.
2. Consider basic choices such as a hash map, min-heap, queue, streaming or channel mechanism, mutex, read-write mutex, and atomic operation.
3. Exclude unused mechanisms, business-unit field definitions, and implementation details that belong in Sections 3.2 or 4.
4. Recommend the relevant basic technical choices to the user and explain briefly where each is used and why it fits.
5. Wait for the user's confirmation of the proposed choices.
6. Generate one sequentially numbered heading for each confirmed choice.
7. Name each heading after the technical mechanism itself, such as `Hash Map`, `Min-Heap`, `Streaming`, `Mutex and RWMutex`, or `Atomic Operations`.
8. Do not use business problems, failure events, or edge cases such as `Duplicate Requests`, `Concurrent Requests`, or `Server Exit` as Section 3.1 headings.
9. For each storage or processing choice, describe its main-process usage, managed data or state, selection reason, key operations or ordering rules, and update invariants.
10. For each lock or atomic choice, describe the protected state, lock mode or atomic boundary, acquisition and release rules, and interaction with other selected mechanisms.
11. Keep all choices consistent with the main process and component workflows in Section 2, the business-unit structures in Section 3.2, and the code in Section 4.

Section 3.1 has no fixed required-topic or optional-topic headings. Generate only confirmed basic technical-choice headings, without placeholders or empty subsections.

### 6.4 Section 3.2: Business Unit Storage Structures

- Identify business units relevant to the current project.
- Discuss which business units need independent storage structures.
- Generate one sequentially numbered, business-specific heading for each relevant business unit structure.
- Replace the business-unit placeholder with the actual structure name, such as `Server Structure`.
- Do not create placeholder business units that are not part of the project.
- Do not generate storage structures for business units that do not need independent storage.
- Define every business unit structure in a fenced `go` code block.
- Before each related type, define only the constants required for business states, fixed limits, default values, or reused fixed values.
- Use typed constants or `iota` when they make the state model clearer and type-safe.
- Do not invent constants that are not required by the business logic.
- Add a concise end-of-line Mandarin comment to every field that carries business meaning, state, a counter, an index, a flag, a limit, a timestamp, or a relationship.
- Make each field comment explain the field's business purpose; do not merely translate or repeat the field name.
- Keep identifiers, field types, constant values, and relationships consistent with Sections 1, 2, and 3.1.
- Keep the Go definitions minimal and syntactically consistent.
- Do not add methods, workflows, or implementation logic to Section 3.2.

### 6.5 Section 4: Coding

Follow the coding requirements in Section 4.0 of the template.

These rules are internal instructions and MUST NOT be copied into the project document.

Requirements:

- Use Go.
- Implement only the core functions required by the current task.
- Keep supporting functions minimal.
- Define required data structures and function signatures before the main function.
- Show only the main processing flow in the main function.
- Use TODO comments instead of detailed unfinished logic.
- Immediately above every function or method declaration, add a bilingual comment block in this order:
  1. English function summary beginning with the function name.
  2. Corresponding Mandarin function summary beginning with the function name.
  3. `Input:` comment explaining every explicit input parameter.
  4. Corresponding `输入：` comment.
  5. `Output:` comment explaining every return value, including errors.
  6. Corresponding `输出：` comment.
- State `Input: none.` and `Output: none.` with their Mandarin equivalents when applicable.
- Make function comments describe purpose and business behavior, not merely restate the signature.
- Add paired English and Mandarin comments immediately before every key step inside a function.
- Treat validation, state transitions, storage updates, lock operations, external calls, error paths, and other business-significant operations as key steps.
- Make step comments explain why the step exists and what business effect it has.
- Add a concise bilingual end-of-line comment to every data-structure field that carries business meaning, with English before Mandarin.
- When concurrency applies, add paired English and Mandarin locking or atomicity rules before the affected function or code block.
- Document lock order, locks that must not be held together, release requirements before network calls, callbacks, waits, or returns, and atomic boundaries only when the confirmed design requires them.
- Do not invent locking rules or comments that are unsupported by the design.
- Keep function signatures and calls consistent.

For Section 4.1:

- Generate `4.1 Main Process Code` only when requested.
- Generate `4.1.1 Main Function` and `4.1.2 Sub-Function Definitions` only when needed.
- Keep the main function focused on the main processing flow.

For Section 4.2:

- Add code only for business units required by the project.
- Generate business-unit headings only for relevant business units.
- Show only each business unit's main processing flow.
- Do not add detailed implementation logic to business-unit processes.
- Add paired English and Mandarin comments before each key step and then call its supporting function.

For Section 4.3:

- Include error and concurrency handling when relevant to the implementation.
- Keep the discussion consistent with the design and storage model.

### 6.6 Section 5: Test Data

Before generating test scenarios:

1. Identify concrete scenarios that cover all required test categories.
2. Analyze the selected difficult problems and edge cases in Section 1.4.
3. Consider the business logic in Section 2, storage design in Section 3, code in Section 4, concurrency, and project-specific risks.
4. Recommend relevant risk-specific scenarios.
5. Present the proposed concrete scenario list to the user and wait for confirmation.
6. Generate required-coverage scenarios first, followed by confirmed risk-specific scenarios, using one continuous numbering sequence.
7. Give every scenario a bilingual heading that names the actual problem or behavior, such as `Invalid Input` / `非法输入` or `Concurrent Requests` / `并发请求`.
8. Never use generic coverage categories such as `Functional Requirements`, `Validation`, or `Concurrent Operations` as final test headings.
9. Generate `Input` and `Expected Response` child headings for every scenario.
10. Encode both Input and Expected Response as Go test data in fenced `go` code blocks.

Required coverage categories:

- Functional Requirements
- Non-Functional Requirements
- Validation
- Core Entity and State Lifecycle
- Failure Handling

Risk-specific coverage categories to consider when relevant:

- Duplicate and Idempotent Operations
- Concurrent Operations
- Resource Availability
- Lock Order and Deadlock Prevention
- Recovery
- Storage Consistency
- Context Management
- Streaming
- External Service Integration

The coverage-category names are internal selection criteria. Do not output them unless one is also the precise name of the actual tested problem.

For every Input and Expected Response Go code block:

- Use composite literals, slices, maps, constants, variables, or error values that match the types and identifiers defined in Sections 3 and 4.
- Prefer concise test-data declarations over full test functions unless the user explicitly requests executable tests.
- List all required business fields first in each struct literal or ordered test-data definition.
- Append included optional business fields after required fields, in business-defined order.
- Omit optional fields that are irrelevant to the scenario.
- Add concise bilingual Go comments to important fields and values, with English first and Mandarin second.
- Make comments explain the field's role or significance in the test rather than repeat its name.
- Keep field ordering deliberate and consistent across related scenarios.
- Keep Input and Expected Response consistent with Sections 1 through 4.

## 7. Cross-Section Consistency

1. Section 1.4 defines the selected difficult problems and edge cases.
2. Section 5 must include corresponding test coverage for relevant selected topics.
3. Risk-specific scenarios in Section 5 must also consider business logic, storage design, concurrency, and project-specific risks.
4. Unselected optional topics must not appear in the final document.
5. Required test topics in Section 5 must always be preserved.
6. When an optional topic is added or removed, update all corresponding sections.
7. Changes to shared entities, states, terminology, or business rules must be reflected in all affected sections.
8. Sections 1 through 5 must describe the same system behavior and architecture.
9. Do not generate dependent headings before their corresponding design decisions have been confirmed.
10. Do not modify unrelated sections when updating cross-section consistency.

## 8. Before Editing

1. Read the bundled template completely.
2. Read and follow Section 0.0.
3. Reload the current target file because the user may have edited it.
4. Identify which sections and headings already exist.
5. Preserve existing user content outside the requested section.
6. Generate only the headings and content required by the current request.
7. Never overwrite an existing system design file without the user's explicit request.

## 9. Start a New System Design

When the user asks to start, build, or create a system design:

1. Read the bundled template completely.
2. Internalize Section 0.0.
3. Create a new Markdown file.
4. Generate only the project title and outermost design headings.
5. Exclude Section 0.0 and all internal instructions.
6. Do not generate second-level or deeper headings.
7. Do not generate required or optional topic lists.
8. Replace the project-name placeholder if known.
9. Do not add project-specific design details yet.
10. Follow the progressive document construction workflow for subsequent requests.

Use a path supplied by the user. Otherwise use `SYSTEM_DESIGN.md` in the current workspace.

If that file already belongs to another project, use `<PROJECT_NAME>_SYSTEM_DESIGN.md` and do not overwrite the existing file.

If the user explicitly asks to complete the entire design in one request, follow all section-specific discussion and confirmation requirements. Generate the complete applicable hierarchy and content only after the necessary confirmations.

## 10. Final Verification

Before completing a design update, verify:

- [ ] The correct project design file was modified.
- [ ] Only requested sections were changed.
- [ ] No unrequested deep-level headings were generated.
- [ ] All applicable rules in Section 0.0 have been followed.
- [ ] Section 0.0 and its rules do not appear in the project document.
- [ ] English content appears before its corresponding Mandarin content.
- [ ] All generated headings use indexed numbering.
- [ ] All bracketed placeholders in completed sections have been replaced.
- [ ] No internal rules or instructional paragraphs appear in the final document.
- [ ] Organizational topic-selection headings do not appear in the final document.
- [ ] All required topics are preserved.
- [ ] Only confirmed optional topics are included.
- [ ] Section 1.4 and Section 5 have consistent topic coverage.
- [ ] Every Section 1.4 child heading names an actual problem, conflict, failure event, or edge case rather than a category or definition.
- [ ] Section 3.1 contains only confirmed foundational storage, data-processing, synchronization, and atomic-operation choices used by the main process.
- [ ] Every Section 3.1 heading names the basic technical choice itself and does not describe a business problem or edge case.
- [ ] Every Section 3.2 business unit structure is defined in a fenced Go code block.
- [ ] Required business constants are defined before their related types, and no unnecessary constants were invented.
- [ ] Every business-meaningful field in Section 3.2 has a concise Mandarin end-of-line comment.
- [ ] Every Section 4 function and method has bilingual purpose, input, and output comments immediately above its declaration.
- [ ] Every key step inside Section 4 functions has a useful paired English and Mandarin comment.
- [ ] Every business-meaningful field defined in Section 4 has a bilingual end-of-line comment.
- [ ] Applicable locking and atomicity rules are documented in bilingual comments without inventing unsupported constraints.
- [ ] Every Section 5 heading names the actual tested problem or behavior rather than a generic coverage category.
- [ ] Every test scenario contains Input and Expected Response as Go test data in fenced `go` code blocks.
- [ ] Required test-data fields appear first and included optional fields follow them in business order.
- [ ] Important test-data fields and values have useful bilingual Go comments.
- [ ] Storage, business logic, code, and tests are consistent.
- [ ] Existing user content has been preserved.

## 11. Boundaries

- Do not modify the bundled template when creating a project design.
- Do not fill unrequested sections during the normal section-by-section workflow.
- Do not generate the complete heading hierarchy during initial document creation.
- Do not generate required or optional topic lists before the corresponding discussion and confirmation.
- Never output Section 0.0.
- Never ignore Section 0.0 when generating any content.
- Do not output internal `.0 Rule` headings or their instructions.
- Do not output organizational topic-selection headings.
- Do not overwrite an existing system design file without the user's explicit request.
- Do not apply this skill to modifications of an existing project or established system design.
- Do not output the complete template specification as the final system design document.
- Do not treat internal authoring instructions as project requirements.
- Do not confuse document structure generation with design content completion.
