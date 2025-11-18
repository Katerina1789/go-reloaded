# task_list.md

## Agile Task Breakdown — `go-reloaded`

This document contains a complete list of implementation tasks for building the `go-reloaded` text transformation tool. Each task follows a TDD-first approach and is designed for incremental development using FSM architecture.

---

### TASK-001: File I/O Setup

- **Functionality**: Read input file and write to output file
- **Test First**: Verify that content is read and written unchanged
- **Implementation Goal**: Build CLI using `os.Args`, `os.ReadFile`, `os.WriteFile`
- **Validation**: Output file matches input when no rules are applied

---

### TASK-002: Tokenization Engine

- **Functionality**: Split text into tokens, combining range patterns
- **Test First**: `"hello (up, 2)"` → `["hello", "(up, 2)"]`
- **Implementation Goal**: Create tokenizer that combines range patterns like `(up, 2)` into single tokens
- **Validation**: Range patterns properly combined, other tokens separated

---

### TASK-003: FSM Controller

- **Functionality**: Manage state transitions and rule dispatching
- **Test First**: `"hello (up)"` routes to uppercase handler
- **Implementation Goal**: Implement FSM with states `Normal` and `QuoteOpen`
- **Validation**: FSM transitions and rule routing work correctly

---

### TASK-004: Rule Handler — Hexadecimal

- **Functionality**: Convert previous word from hex to decimal
- **Test First**: `"1E (hex)"` → `"30"`
- **Implementation Goal**: Implement hex parsing and replacement
- **Validation**: Unit tests for valid/invalid hex inputs

---

### TASK-005: Rule Handler — Binary

- **Functionality**: Convert previous word from binary to decimal
- **Test First**: `"101 (bin)"` → `"5"`
- **Implementation Goal**: Implement binary parsing and replacement
- **Validation**: Unit tests for valid/invalid binary inputs

---

### TASK-006: Rule Handler — Casing (Single)

- **Functionality**: Transform previous word casing
- **Test First**: `"hello (up)"` → `"HELLO"`
- **Implementation Goal**: Implement `(up)`, `(low)`, `(cap)`
- **Validation**: Unit tests for each rule

---

### TASK-007: Rule Handler — Casing (Range)

- **Functionality**: Transform previous `n` words casing
- **Test First**: `"this is exciting (up, 3)"` → `"THIS IS EXCITING"`
- **Implementation Goal**: Extend casing logic to handle ranges
- **Validation**: Unit tests for edge cases

---

### TASK-008: Rule Handler — Punctuation

- **Functionality**: Normalize punctuation spacing
- **Test First**: `"hello ,world"` → `"hello, world"`
- **Implementation Goal**: Attach punctuation to previous word, space from next
- **Validation**: Unit tests for all punctuation types

---

### TASK-009: Rule Handler — Quotes

- **Functionality**: Format quoted phrases
- **Test First**: `" ' awesome ' "` → `"'awesome'"`
- **Implementation Goal**: Detect quote boundaries and trim spacing
- **Validation**: Unit tests for single and multi-word quotes

---

### TASK-010: Post-Processing Pipeline

- **Functionality**: Apply punctuation, quote, and article fixes after FSM processing
- **Test First**: `"hello ,world ' text '"` → `"hello, world 'text'"`
- **Implementation Goal**: Implement `FixPunctuation()`, `FixQuotes()`, `FixArticles()` functions
- **Validation**: Unit tests for each post-processing function

---

### TASK-011: Integration Test — Full Paragraph

- **Functionality**: Build comprehensive integration tests that validate the entire text transformation pipeline
- **Test First**: Golden test framework with 5 test cases of increasing complexity
- **Implementation Goal**: End-to-end pipeline validation with performance benchmarks
- **Validation**: Byte-for-byte output matching, <100ms processing time for 1000-word paragraphs

---

### TASK-012: CLI & Error Handling

- **Functionality**: Comprehensive error handling with user-friendly messages and robust validation
- **Test First**: All error scenarios - missing files, permission issues, invalid arguments
- **Implementation Goal**: Production-ready CLI with clear feedback and Unix exit codes
- **Validation**: Clear error messages, proper exit codes (0=success, 1=error, 2=usage)

---

### TASK-013: Audit Mode

- **Functionality**: Automated validation system with comprehensive test suite and detailed reporting
- **Test First**: 12 test cases covering all rules and edge cases with pass/fail checklist
- **Implementation Goal**: Self-contained audit mode with --audit flag, detailed reports
- **Validation**: Complete rule coverage, <5s execution time, actionable failure reports

---

## Completion Criteria

Each task is complete when:

- All related tests pass
- Output matches expected results
- Code follows Go best practices
- Audit checklist is satisfied
