---
name: code-implementation-workflow
description: Implement a confirmed system design as production-ready Go code in an existing workspace, organize source and tests by business unit, run the test suite, and produce a factual bilingual test report and README. Use when the user asks to finish or build code from a design and verify it; do not use for design-only requests, isolated explanations, or reviews without implementation.
---

# Code Implementation Workflow

Turn an approved design into complete, verified Go code. Work in the user's current project unless the user provides another target.

## 1. Scope and Authorization

- Treat implementation requests as authorization to edit source files, test files, and the local test report inside the selected workspace.
- Do not deploy, publish, push, send external messages, or modify external systems unless the user separately requests it.
- Do not replace confirmed architecture or business rules with a materially different design. Ask for direction when a missing decision would change public behavior, storage semantics, or compatibility.
- Make reasonable, local implementation choices without stopping for minor details.

## 2. Inspect Before Editing

1. Read applicable repository instructions such as `AGENTS.md`.
2. Search for the applicable System Design file, including names such as `SYSTEM_DESIGN.md` or `*_SYSTEM_DESIGN.md`, and read it completely when present.
3. Reload the confirmed design, requirements, existing source, tests, and current worktree state.
4. Identify the Go module, package layout, entry points, build commands, and existing conventions.
5. Map each business unit and the main process to a focused source-file or package boundary.
6. Identify the required tests from the design, especially difficult problems, edge cases, workflows, storage rules, and concurrency guarantees.
7. Preserve unrelated user changes and reuse compatible existing code.

When an applicable System Design file exists, treat it as the implementation source of truth for business units, workflows, structures, storage choices, synchronization rules, function contracts, and test scenarios. Do not silently diverge from it. If the user's current request materially conflicts with the design, explain the conflict and obtain direction before implementing the conflicting behavior.

If the repository has no usable design document, derive the minimum implementation plan from the user's requirements and existing code. Ask only when a material behavior or architecture choice cannot be inferred safely.

## 3. Organize Code by Business Unit

- Do not put the complete implementation in one source file.
- Separate cohesive business units into focused files or packages according to the repository's existing structure.
- Keep the main entry point focused on orchestration. Move business rules, state changes, storage operations, and synchronization into their owning units.
- Put genuinely shared types or helpers in focused shared files. Do not create a generic utility file for unrelated functions.
- Avoid artificial fragmentation: tightly related constants, types, and operations may stay together.
- Mirror source boundaries in `_test.go` files. Do not put all unit tests in one large test file.

A small project may use a layout such as:

```text
main.go
request.go
request_test.go
server.go
server_test.go
load_balancer.go
load_balancer_test.go
```

Adapt the layout to the actual business units; do not copy this example mechanically.

## 4. Implement the Confirmed Design

- Use Go and keep identifiers, constants, structures, states, errors, function signatures, and workflows consistent with the confirmed design.
- Implement all behavior in the requested scope. Do not leave placeholder bodies, fake success paths, or unresolved TODOs for required logic.
- Apply the selected storage and synchronization mechanisms consistently. Preserve lock ownership, atomic boundaries, update ordering, cleanup, and rollback rules.
- Return explicit errors for invalid input and failed operations. Do not silently discard errors required by the design.
- Keep interfaces and supporting functions as small as the design permits.
- Follow the repository's compatibility and dependency conventions. Add or upgrade dependencies only when the implementation genuinely requires it.

Use the following code-comment contract:

- Explain every business constant. Put a concise English explanation and corresponding Mandarin explanation directly above the constant or constant group. When grouped constants have different meanings, explain each constant separately.
- Add a concise English summary and corresponding Mandarin summary immediately above every function or method.
- Explain every explicit input parameter and every output value, including errors, directly above the function or method declaration. Describe receiver behavior in the method summary when it affects state.
- State explicitly when a function has no input or no output.
- Add paired English and Mandarin comments before business-significant validation, state changes, storage updates, lock operations, external calls, and error paths.
- Explain every business-significant `if`, `switch`, retry, fallback, or early-return branch in the processing flow. State why the branch is taken and what business result it produces; do not merely restate the condition.
- Add useful comments to fields whose business meaning is not obvious.
- Document only concurrency rules that the design actually requires.

For an established repository with a different documented comment style, preserve its formatting where possible, but do not omit constant meanings, function purpose, input and output explanations, or business-significant branch-flow explanations.

## 5. Write Runnable Unit Tests

- Create real Go unit tests in `_test.go` files. Do not provide pseudocode or documentation-only test data as a substitute.
- Immediately above every test function, add an English explanation and corresponding Mandarin explanation of the concrete situation being tested.
- Make the test-function explanation state the setup or trigger, the behavior under test, and the expected result or guarantee. Do not use a generic description such as "tests ProcessRequest."
- Begin the English test comment with the Go test function name when required by the repository's linting convention.
- Cover the normal path and every applicable confirmed scenario from the design.
- Include validation, entity and state lifecycle, and failure handling. Include duplicate operations, concurrency, resource availability, recovery, storage consistency, streaming, or external integration only when relevant.
- Use table-driven tests when several cases share the same setup and assertion pattern.
- For a table-driven test, make the test-function comment explain the overall scenario group and give every case a concrete `name` that states its individual situation and expected behavior.
- Give every `t.Run` subtest a descriptive scenario name. Add paired English and Mandarin comments for non-obvious case setup or expectations.
- Keep tests deterministic and isolated. Avoid arbitrary sleeps for concurrency tests; prefer channels, wait groups, barriers, contexts, and bounded deadlines.
- Verify observable behavior, state changes, invariants, errors, and cleanup. Do not assert implementation details unless they are part of the design contract.
- Do not weaken, delete, or rewrite valid tests merely to make the suite pass.
- Keep fixtures minimal and place required fields before optional fields when constructing test data.

## 6. Format and Verify

Run the checks supported by the project. For an ordinary Go module, use this sequence:

1. Format changed Go files with `gofmt`.
2. Run focused tests for the changed business units.
3. Run the complete suite with `go test ./...`.
4. Run `go test -race ./...` when concurrency or shared mutable state is relevant and the environment supports the race detector.
5. Run `go vet ./...` when it is available and appropriate for the project.
6. Run coverage only when requested or useful, and report the exact command and result.

Fix failures caused by the implementation and rerun the affected checks. Do not claim a check passed unless its command completed successfully.

If verification is blocked by unavailable dependencies, permissions, platform limitations, or unrelated pre-existing failures:

- Exhaust safe in-scope alternatives.
- Preserve the exact command and relevant failure output.
- Distinguish implementation failures from environmental or pre-existing failures.
- Do not conceal the blocker or fabricate a passing result.

## 7. Generate the Test Report

After running the tests, create a Markdown report at the user-specified path. Otherwise use `TEST_REPORT.md` in the project root, unless the environment requires user-facing artifacts in a dedicated output directory.

If the default report path already contains user-authored material, preserve it and create a clearly named new report instead of overwriting it.

Write the report in bilingual mode:

- Put English first and the corresponding Mandarin translation immediately after it.
- Use bilingual section headings, summaries, test-scenario descriptions, result explanations, failure analysis, limitations, and follow-up notes.
- Use bilingual labels in tables and status summaries, such as `Passed / 通过`, `Failed / 失败`, `Skipped / 跳过`, `Blocked / 受阻`, and `Not run / 未运行`.
- Keep commands, file paths, package names, function names, error identifiers, and raw tool output unchanged. Add bilingual explanations around them instead of translating or rewriting the raw technical data.
- Ensure the English and Mandarin content communicate the same facts; do not add conclusions in only one language.

The report must contain:

1. Implementation scope and business units.
2. Source and test file organization.
3. Commands actually executed.
4. Test scenarios covered.
5. Passed, failed, skipped, and blocked results.
6. Race-detector, vet, and coverage results when those checks were run.
7. Remaining failures, limitations, or known gaps.

Base every result on actual command output. Mark checks that were not run as `Not run`; never infer or invent results.

## 8. Create or Update the README

After the implementation, tests, and bilingual test report are complete, create or update `README.md` in the project root or the repository's established documentation location.

- If a README already exists, preserve unrelated user-authored content. Add or update only the sections needed to describe the completed implementation.
- Write the README in bilingual mode with English first and the corresponding Mandarin text immediately after it.
- Keep the English and Mandarin descriptions factually equivalent.
- Base the README on the confirmed System Design and the final implemented files. Do not describe planned behavior that was not implemented.

The README must include at least:

1. `Requirements / 需求`: summarize the implemented functional requirements and the important non-functional, validation, lifecycle, failure-handling, storage, or concurrency requirements that affect usage or correctness.
2. `File Responsibilities / 文件主要功能`: list every maintained project file created or materially changed by the implementation and explain its main responsibility. Also include other major entry-point, configuration, System Design, test-report, source, and test files needed to understand the project. Omit generated caches, dependency contents, binaries, and temporary artifacts.

Use a bilingual table for file responsibilities when it improves clarity. Keep file paths and identifiers unchanged, and write their explanations in English first and Mandarin second.

Link the applicable System Design and `TEST_REPORT.md` from the README when those files are available.

## 9. Completion Criteria

The task is complete only when:

- The requested implementation is present and consistent with the design.
- Source code and tests are separated by business unit without unnecessary fragmentation.
- The implementation follows the applicable System Design file when one exists.
- Required code contains no incomplete stubs or placeholder success paths.
- Business constants, function purpose, inputs, outputs, and significant processing branches have useful explanations.
- Unit tests are runnable.
- Every test function explains the concrete tested situation and expected result, and every table-driven case has a descriptive scenario name.
- All feasible relevant checks have been executed and implementation-caused failures have been addressed.
- A factual bilingual English-and-Mandarin test report has been created.
- A bilingual README has been created or updated after implementation and testing, and it accurately explains the implemented requirements and the main responsibility of each relevant project file.

In the final response, lead with the implementation and test outcome. Link the main changed files, README, and test report, list the commands run, and state any unresolved blocker or unverified check.
