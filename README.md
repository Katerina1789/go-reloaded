# go-reloaded

A text transformation tool written in Go that applies formatting, correction, and grammar rules to input files. Designed for audit-based evaluation and collaborative testing.

## Overview

`go-reloaded` reads a text file, applies a series of transformation rules (e.g., number conversion, capitalization, punctuation correction), and writes the modified output to a new file. It supports both simple and compound rule interactions, making it suitable for nuanced text editing tasks.
This project is part of a collaborative audit framework where students act as both developers and auditors.

## Workflow Diagram

```mermaid
flowchart TD
    A[Start: Read input file] --> B[Parse text line by line]
    B --> C[Apply transformation rules]
    C --> D{Rules Detected?}
    D -- Yes --> E[Apply matching rules]
    D -- No --> F[Skip to next line]
    E --> G[Update transformed line]
    F --> G
    G --> H[Write to output file]
    H --> I[End]

## Repository Structure

graph TD
    A[go-reloaded] --> B[audit/]
    A --> C[docs/]
    A --> D[README.md]
    B --> E[audit_guide.txt]
    B --> F[golden_test_set.txt]
    C --> G[analysis.md]

-  Analysis Document: Problem description, rule breakdown, and architectural approach (docs/analysis.md)
-  Golden Test Set: Functional and custom test cases in natural language (audit/golden_test.txt)
-  Audit Guide: Checklist for verifying correctness and marking pass/fail outcomes (audit/audit_guide.txt)

## Usage

To run the tool:

**go run . input.txt output.txt**

Example:
cat sample.txt
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
go run . sample.txt result.txt
cat result.txt
Simply add 66 and 2 and you will see the result is 68.

## Testing & Auditing

- All test cases are written in natural language for clarity and accessibility.
- Auditors should use the audit guide to verify outputs and mark results.
- The golden test set includes both audit examples and original edge cases.

## External References

- Go Documentation: https://golang.org/doc/
- Go File System API: https://pkg.go.dev/os
- Markdown/HTML Guide: https://www.markdownguide.org/basic-syntax/
- Mermaid diagrams: https://github.com/mermaid-js/mermaid
- Finite State Machines – CS Concepts: https://en.wikipedia.org/wiki/Finite-state_machine