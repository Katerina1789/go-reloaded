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

```bash
# From project root (go.mod is in config/)
GOMOD=config/go.mod go run . input.txt output.txt
GOMOD=config/go.mod go run . input.txt output.txt --audit
```

## Repository Structure

```
go-reloaded/
├── audit/                  # Audit checklists and test cases
│   ├── audit_guide.md
│   ├── audit_tasks.md
│   └── golden_test_set.md
├── config/                 # Configuration files
│   ├── .gitignore
│   ├── go.mod
│   └── README.md
├── docs/                   # Architecture, agent instructions, and references
│   ├── analysis.md
│   ├── architecture.md
│   ├── how_to_work.md
│   ├── meta_prompt.md
│   └── references.md
├── tasks/                  # Agile task breakdown
│   └── task_list.md
├── agents.md
├── LICENSE
└── README.md               # Project overview and usage
```

## Development Workflow

1. Analyze the problem and architecture
2. Write tests first (TDD)
3. Implement rule handlers and FSM logic
4. Validate using audit checklists and golden test cases
5. Document edge cases and commit changes

See [docs/how_to_work.md](docs/how_to_work.md) for a full step-by-step guide.

## References

- [FSM design and rationale](docs/architecture.md)
- [Rule breakdown and analysis](docs/analysis.md)
- [Audit checklist and test cases](audit/audit_tasks.md)
- [Golden test set](audit/golden_test_set.md)
- [Audit guide](audit/audit_guide.md)
- [AI agent instructions](agents.md)
- [Meta-prompting guide](docs/meta_prompt.md)
- [Development workflow](docs/how_to_work.md)
- [Task breakdown](tasks/task_list.md)

## License

MIT License - see LICENSE file for details