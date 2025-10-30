# Rule Handler — Casing (Range)

- **ID**: TASK-007
- **Owner**: Backend Developer
- **Size**: M
- **Confidence**: Medium
- **Hard Dependencies**: TASK-003, TASK-006
- **Soft Dependencies**: None
- **Related Architecture**: Rule handlers, range processing

## Mission Profile
- Extend casing transformation to handle multiple words with range syntax
- Parse range parameters from rule markers like `(up, 3)`, `(low, 2)`, `(cap, 5)`
- Apply transformations to the specified number of previous words
- Handle edge cases where range exceeds available words

## Deliverables
- Extension to `internal/rules/casing.go` with range processing logic
- Range parameter parsing and validation
- Backward token traversal for multi-word transformation
- Comprehensive unit tests for range operations and edge cases
- Error handling for invalid range parameters

## Acceptance Criteria
- ✅ `"this is exciting (up, 3)"` correctly transforms to `"THIS IS EXCITING"`
- ✅ `"hello world (cap, 2)"` correctly transforms to `"Hello World"`
- ✅ Range exceeding available words handles gracefully (transforms available words only)
- ✅ Invalid range parameters (negative, zero, non-numeric) are handled properly
- ✅ Integration with existing single-word casing logic

## Verification Plan
- `unit`: Test range parsing and multi-word transformations
- `edge`: Test ranges exceeding word count and invalid parameters
- `integration`: Test with FSM controller and token processing
- `regression`: Ensure single-word casing still works correctly

## References
- `docs/architecture.md`: Rule handler interface and token processing
- `audit/golden_test_set.md`: Range casing transformation test cases
- TASK-006 implementation for single-word casing patterns

## Notes for Agent
- Extend existing casing handler rather than creating separate file
- Implement robust range parameter parsing with error handling
- Consider token boundary detection for accurate word counting

## PROMPT — FULL 4-STEP FLOW

```markdown
You are executing **Rule Handler — Casing (Range) (TASK-007)** for go-reloaded.

### Step 1 — Analyze & Confirm
- Review single-word casing implementation from TASK-006
- Study FSM controller token processing from TASK-003
- Examine `audit/golden_test_set.md` for range casing test cases
- WAIT for user confirmation before proceeding

### Step 2 — Write the Tests (TDD)
- Create unit tests for range parameter parsing
- Test multi-word transformations with various ranges
- Test edge cases (range > available words, invalid params)
- Prepare integration tests with FSM controller

### Step 3 — Implement the Code
- Extend casing handler with range processing logic
- Implement backward token traversal for multi-word selection
- Add range parameter validation and error handling
- Integrate with existing FSM controller interface

### Step 4 — Validate & QA
- Run all range casing transformation tests
- Test integration with FSM and tokenizer
- Validate edge case handling and error conditions
- Ensure backward compatibility with single-word casing
- If verification passes, output: **"✅ Rule Handler — Casing (Range) (TASK-007) self-verified. Ready for review."**
```