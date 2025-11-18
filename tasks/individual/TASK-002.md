# Tokenization Engine

- **ID**: TASK-002
- **Owner**: Backend Developer
- **Size**: S
- **Confidence**: High
- **Hard Dependencies**: TASK-001
- **Soft Dependencies**: None
- **Related Architecture**: Text processing, FSM preparation

## Mission Profile
Split text into tokens, combining range patterns like `(up, 2)` into single tokens. Create tokenizer that preserves punctuation and rule markers while handling special range pattern combinations.

## Deliverables
- Tokenization function in `main.go`
- Logic to combine range patterns `(up, 2)`, `(low, 3)`, `(cap, 4)` into single tokens
- Proper handling of spaces and word boundaries
- Unit tests for tokenization edge cases

## Acceptance Criteria
- ✅ `"hello (up, 2)"` → `["hello", "(up, 2)"]`
- ✅ Range patterns properly combined into single tokens
- ✅ Regular words separated correctly
- ✅ Handles edge cases: multiple spaces, punctuation
- ✅ Preserves original text structure for reconstruction

## Verification Plan
- `unit`: Test range pattern combining, word separation
- `integration`: Test with complex text containing multiple patterns
- `edge`: Test empty strings, single words, multiple spaces
- `performance`: Verify efficient tokenization for large texts

## References
- Go `strings` package documentation
- `docs/analysis.md`: Rule specifications
- FSM requirements for token processing

## PROMPT — FULL 4-STEP FLOW

```markdown
You are executing **Tokenization Engine (TASK-002)** for go-reloaded.

### Step 1 — Analyze & Confirm
- Review range pattern requirements: `(up, 2)`, `(low, 3)`, `(cap, 4)`
- Understand tokenization needs for FSM processing
- WAIT for user confirmation before proceeding

### Step 2 — Write the Tests (TDD)
- Test range pattern combining
- Test word separation and boundary handling
- Test edge cases and malformed input

### Step 3 — Implement the Code
- Build tokenization function with range pattern logic
- Implement proper space and punctuation handling
- Ensure tokens can be rejoined correctly

### Step 4 — Validate & QA
- Run all tokenization tests
- Verify range patterns combine correctly
- Test with sample text files
- If verification passes, output: **"✅ Tokenization Engine (TASK-002) self-verified. Ready for review."**
```