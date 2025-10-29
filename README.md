# go-reloaded

## Overview

`go-reloaded` is a rule-based text transformation tool written in Go. It reads an input file, applies a series of formatting and grammar rules, and writes the corrected output to a new file. The project is designed using a Finite State Machine (FSM) architecture and follows Test-Driven Development (TDD) principles.

This repository is structured to support audit-driven development and AI-assisted implementation. It includes detailed documentation, audit checklists, and a modular task system for incremental development.

---

## Features

- Rule-based transformations:
  - Casing: `(up)`, `(low)`, `(cap)` and range variants `(up, n)`, `(low, n)`, `(cap, n)`
  - Number conversion: `(hex)`, `(bin)`
  - Article correction: `a` → `an` before vowels or 'h'
  - Punctuation normalization
  - Quote formatting using `' '`
- FSM-based architecture for context-aware rule application
- Audit-ready test cases and verification guides
- AI agent support via meta-prompting and structured task flow

---

## Getting Started

## Run the tool

go run . input.txt output.txt
go run . input.txt output.txt --audit

## Repository Structure

go-reloaded/
├── audit/                  # Audit checklists and test cases
│   ├── audit_guide.txt
│   ├── audit_tasks.md
│   ├── golden_test_set.txt
│   └── sample.txt
├── docs/                   # Architecture, agent instructions, and references
│   ├── analysis.md
│   ├── architecture.md
│   ├── how_to_work.md
│   ├── agents.md
│   ├── meta_prompt.md
│   └── references.md
├── tasks/                  # Agile task breakdown
│   └── task_list.md
└── README.md               # Project overview and usage

## Development Workflow

Analyze the problem and architecture.

Write tests first (TDD).

Implement rule handlers and FSM logic.

Validate using audit checklists and golden test cases.

Document edge cases and commit changes.

See docs/how_to_work.md for a full step-by-step guide.

## References

FSM design and rationale: docs/architecture.md

Rule breakdown and analysis: docs/analysis.md

Audit checklist and test cases: audit/audit_guide.txt, audit/golden_test_set.txt

AI agent instructions: docs/agents.md, docs/meta_prompt.md

External resources and further reading: docs/references.md

## License