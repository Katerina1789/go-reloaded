# Rule Handler — Casing (Single)

- **ID**: TASK-006
- **Owner**: Backend Developer
- **Size**: S
- **Confidence**: High
- **Hard Dependencies**: TASK-003
- **Soft Dependencies**: None
- **Related Architecture**: Rule handlers, text transformation

## Mission Profile
- Implement single-word casing transformation rules
- Handle uppercase `(up)`, lowercase `(low)`, and capitalize `(cap)` operations
- Apply transformations to the previous word in the token stream
- Integrate with FSM controller for rule dispatching

## Deliverables
- `internal/rules/casing.go` with single-word casing logic
- Support for `(up)`, `(low)`, and `(cap)` rule markers
- Integration with FSM rule routing system
- Comprehensive unit tests for all casing operations
- Unicode-aware text transformation handling

## Acceptance Criteria
- ✅ `"hello (up)"` correctly transforms to `"HELLO"`
- ✅ `"WORLD (low)"` correctly transforms to `"world"`
- ✅ `"example (cap)"` correctly transforms to `"Example"`
- ✅ Handles Unicode characters and special cases properly
- ✅ Integration with FSM controller works seamlessly

## Verification Plan
- `unit`: Test each casing rule with various input strings
- `unicode`: Test with non-ASCII characters and special cases
- `integration`: Test with FSM controller and token processing
- `edge`: Test with empty strings and single characters

## References
- `docs/architecture.md`: Rule handler interface and FSM integration
- `audit/golden_test_set.md`: Casing transformation test cases
- Go `strings` package documentation for case conversion

## Notes for Agent
- Use Go's built-in `strings.ToUpper`, `strings.ToLower`, `strings.Title` functions
- Implement consistent interface pattern with other rule handlers
- Consider Unicode normalization for international text

## PROMPT — FULL 4-STEP FLOW

```markdown
You are executing **Rule Handler — Casing (Single) (TASK-006)** for go-reloaded.

### Step 1 — Analyze & Confirm
- Review FSM controller interface from TASK-003
- Study rule handler patterns from previous tasks for consistency
- Examine `audit/golden_test_set.md` for casing transformation test cases
- WAIT for user confirmation before proceeding

### Step 2 — Write the Tests (TDD)
- Create unit tests for each casing rule (up, low, cap)
- Test Unicode character handling
- Test integration with FSM rule routing
- Prepare edge case tests (empty strings, single chars)

### Step 3 — Implement the Code
- Build casing transformation handlers using Go's strings package
- Implement rule detection and application logic
- Integrate with FSM controller interface
- Add support for Unicode text processing

### Step 4 — Validate & QA
- Run all casing transformation tests
- Test integration with FSM and tokenizer
- Validate Unicode handling and edge cases
- Check consistency with other rule handlers
- If verification passes, output: **"✅ Rule Handler — Casing (Single) (TASK-006) self-verified. Ready for review."**
```