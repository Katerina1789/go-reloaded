# how_to_work.md

## How to Work in the `go-reloaded` Repository

This document provides a step-by-step guide for AI agents and junior developers contributing to the `go-reloaded` project. The goal is to build a rule-based text transformation tool in Go using Test-Driven Development (TDD) and FSM architecture.

---

## Prerequisites

Before starting, make sure you:

- Understand the problem domain described in `docs/analysis.md`
- Review the FSM architecture in `docs/architecture.md`
- Familiarize yourself with the audit framework in `audit/`
- Read the task breakdown in `tasks/task_list.md`
- Follow AI agent guidelines in `agents.md` for consistent behavior

---

## Step-by-Step Workflow

### Step 1 — Analyze the Task

- Open `tasks/task_list.md`
- Select the next available task in order
- Read the task description carefully
- Identify any dependencies or required context
- Refer to `docs/analysis.md` and `docs/architecture.md` for rule logic and FSM design

### Step 2 — Write the Tests (TDD)

- Begin by writing unit tests for the described functionality
- Use Go’s `testing` package
- Reference examples from `audit/golden_test_set.md`
- Cover edge cases and tricky rule interactions

### Step 3 — Implement the Code

- Write clean, idiomatic Go code using only standard packages
- Respect FSM transitions and state logic
- Implement rule handlers as modular functions
- Implement helper functions as needed within main.go or separate files
- Keep code readable and maintainable

### Step 4 — Validate the Output

- Run all tests and confirm they pass
- Compare output against expected results in `audit/golden_test_set.md`
- Use `audit/audit_guide.md` to verify formatting, punctuation, and correctness
- Ensure the transformation matches audit expectations

### Step 5 — Document and Commit

- Document any edge cases or implementation notes
- Commit your changes with a clear message
- If operating as an agent, wait for operator confirmation before proceeding

---

## Additional Notes

- Always follow TDD: write tests before code
- Use FSM architecture for rule coordination
- Avoid external dependencies—use only Go standard library
- Ask for clarification if a rule or task is ambiguous
- Follow agent guidelines in `agents.md` for AI-assisted development

---

## Completion Criteria

A task is considered complete when:

- All related tests pass
- Output matches expected results
- Code follows Go best practices
- Audit checklist is satisfied

---

This guide ensures consistent, high-quality contributions to the `go-reloaded` project. Follow each step carefully and refer to the documentation as needed.
