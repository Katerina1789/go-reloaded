# architecture.md

## Architecture Overview for `go-reloaded`

This document outlines the architectural design and rationale for the `go-reloaded` text transformation tool. The project is implemented in Go and follows a Finite State Machine (FSM) model to handle rule-based text editing with precision and contextual awareness.

---

## Problem Domain

The tool processes an input text file and applies a series of transformation rules, including:

- Number conversions: `(hex)`, `(bin)`
- Casing transformations: `(up)`, `(low)`, `(cap)` and their range variants `(up, n)`, `(low, n)`, `(cap, n)`
- Article correction: `a` → `an` before vowels or 'h'
- Punctuation normalization
- Quote formatting using `' '`

The output is a cleaned and grammatically improved version of the original text.

---

## Architectural Decision: FSM vs Pipeline

### Pipeline Model (Rejected)

- Applies rules in a fixed linear sequence
- Each rule operates independently
- Simple to implement but lacks contextual awareness
- Cannot track state (e.g., inside quotes or punctuation boundaries)
- Risk of rule conflicts and incorrect transformations

### FSM Model (Chosen)

- Tracks the current state of the text (e.g., normal, inside quotes, after punctuation)
- Applies rules dynamically based on context
- Handles overlapping and grammar-sensitive rules accurately
- Supports complex interactions between rules
- Requires state definitions and transitions, but offers greater control

---

## FSM Design

### States

- `Normal`: Default state for reading and transforming tokens
- `QuoteOpen`: Inside a quoted phrase
- `RulePending`: Awaiting rule application (e.g., after `(up)`)
- `RangePending`: Awaiting range-based transformation (e.g., `(low, 3)`)
- `PunctuationPending`: Handling punctuation placement
- `ArticleCheck`: Evaluating `a/an` usage

### Transitions

- `Normal` → `RulePending`: When a rule marker like `(up)` is detected
- `RulePending` → `Normal`: After applying the rule
- `Normal` → `QuoteOpen`: When `'` is detected
- `QuoteOpen` → `Normal`: When closing `'` is found
- `Normal` → `ArticleCheck`: When `a` is followed by a vowel-starting word

---

## Rule Handlers

Each rule is implemented as a handler function that:

- Accepts the current token and context
- Applies the transformation
- Updates the token stream
- Returns control to the FSM

### Examples

- `handleHex()`: Converts hexadecimal to decimal
- `handleBin()`: Converts binary to decimal
- `handleUp(n)`: Converts previous `n` words to uppercase
- `handleLow(n)`: Converts previous `n` words to lowercase
- `handleCap(n)`: Capitalizes the first letter of previous `n` words
- `handleQuotes()`: Trims and wraps quoted phrases
- `handlePunctuation()`: Normalizes spacing around punctuation
- `handleArticleCorrection()`: Replaces `a` with `an` when appropriate

---

## Modular Design

- `main.go`: Entry point and CLI handler
- `fsm.go`: FSM controller and state manager
- `rules/`: Contains individual rule handlers
- `utils/`: Tokenization, string helpers, and file I/O
- `tests/`: Unit tests for each rule and integration tests for full paragraphs

---

## Benefits of FSM Architecture

- Context-aware transformations
- Cleaner separation of concerns
- Easier debugging and auditing
- Scalable for future rule additions
- Ideal for audit-driven environments

---

## Testing Framework

### Integration Testing (TASK-011)
- Golden test framework with byte-for-byte output validation
- 5 test cases of increasing complexity (100-5000 words)
- Performance benchmarks: <100ms for 1000-word paragraphs
- Memory profiling and leak detection

### Audit Mode (TASK-013)
- Automated validation system with `--audit` flag
- 20+ comprehensive test cases covering all rules
- Detailed pass/fail reporting with failure diagnostics
- CI/CD integration for regression testing
- Performance requirement: <5 seconds full suite execution

## Next Steps

- Finalize FSM state definitions and transitions
- Implement rule handlers incrementally using TDD
- Validate outputs against `audit/golden_test_set.txt`
- Document edge cases and tricky interactions
- Integrate audit mode into development workflow
