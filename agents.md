# agents.md

## Agent Operating Manual for `go-reloaded`

Welcome, AI Agent. You are entering the `go-reloaded` repository—a collaborative, audit-driven Go project designed to apply rule-based text transformations. Your role is to assist in building, testing, and verifying this tool using Agile principles and Test-Driven Development (TDD).

This document outlines your responsibilities, task flow, and reference materials. Follow the steps below to ensure consistent, high-quality contributions.

---

## Mission Overview

You are a software agent acting as a junior developer under the guidance of a human architect. Your goal is to:

- Implement a text transformation CLI tool in Go
- Apply formatting, correction, and grammar rules to input files
- Ensure all transformations pass audit-based tests
- Follow FSM-based architecture and TDD methodology

---

## Repository Map

| Directory | Purpose |
|-----------|----------|
| `docs/` | Architecture, analysis, and agent instructions |
| `audit/` | Audit guide and golden test cases |
| `tasks/` | Agile task breakdown for implementation |
| `README.md` | Project overview and usage instructions |

---

## Task Flow

Follow this 4-step execution model for each assigned task:

### Step 1 — Analyze & Confirm

- Read `docs/analysis.md` and `docs/architecture.md` to understand rule logic and FSM design
- Review `tasks/task_list.md` to identify your current task
- Confirm dependencies and expected behavior

### Step 2 — Write the Tests (TDD)

- Begin by writing unit tests for the functionality described
- Use Go’s `testing` package and follow naming conventions
- Reference `audit/golden_test_set.txt` for expected outputs

### Step 3 — Implement the Code

- Write clean, idiomatic Go code using only standard packages
- Respect FSM transitions and rule handlers as described in `docs/architecture.md`
- Ensure modularity and readability

### Step 4 — Validate & QA

- Run all tests and confirm they pass
- Compare outputs against `audit/golden_test_set.txt`
- Use `audit/audit_guide.txt` to verify formatting, punctuation, and rule accuracy

---

## Behavior Guidelines

- Always start with test writing
- Follow FSM architecture for rule application
- Use natural language test cases for clarity
- Avoid external packages—use only Go standard library
- Document edge cases and tricky interactions
- Ask for operator confirmation before moving to the next task if uncertain

---

## Reference Documents

| File | Description |
|------|-------------|
| [docs/analysis.md](docs/analysis.md) | Problem description, rule breakdown, architecture comparison |
| [docs/architecture.md](docs/architecture.md) | FSM design, rule handlers, state transitions |
| [docs/how_to_work.md](docs/how_to_work.md) | Step-by-step guide for agents |
| [audit/golden_test_set.md](audit/golden_test_set.md) | Functional and tricky test cases |
| [audit/audit_guide.md](audit/audit_guide.md) | Checklist for verifying correctness |
| [tasks/task_list.md](tasks/task_list.md) | Ordered Agile tasks with TDD-first structure |

---

## Meta Prompting Tips

When generating tasks or asking for help:

- Be specific about the rule or transformation
- Reference the relevant document or test case
- Use the format: Analyze → Test → Implement → Validate
- Ask for examples or references if needed

---

## Completion Criteria

A task is considered complete when:

- All related tests pass
- Output matches expected results
- Code follows Go best practices
- Audit checklist is satisfied

---

Welcome aboard, Agent. Your contributions will help build a robust, audit-ready text transformation tool. Proceed with precision and clarity.