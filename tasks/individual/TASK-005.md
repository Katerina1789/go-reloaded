# Rule Handler — Binary

- **ID**: TASK-005
- **Owner**: Backend Developer
- **Size**: S
- **Confidence**: High
- **Hard Dependencies**: TASK-003
- **Soft Dependencies**: TASK-004 (similar pattern)
- **Related Architecture**: Rule handlers, number conversion

## Mission Profile
- Implement binary to decimal conversion rule handler
- Parse binary strings and convert to decimal representation
- Handle invalid binary inputs gracefully with error reporting
- Integrate with FSM controller for rule dispatching

## Deliverables
- `internal/rules/binary.go` with binary conversion logic
- Binary validation and error handling
- Integration with FSM rule routing system
- Comprehensive unit tests for valid and invalid inputs
- Performance optimization for large binary numbers

## Acceptance Criteria
- ✅ `"101 (bin)"` correctly converts to `"5"`
- ✅ `"1010 (bin)"` correctly converts to `"10"`
- ✅ Invalid binary strings (containing 2-9, letters) are handled gracefully
- ✅ Empty or malformed binary inputs return appropriate errors
- ✅ Integration with FSM controller works seamlessly

## Verification Plan
- `unit`: Test binary parsing with valid inputs (0, 1, 101, 1111, etc.)
- `error`: Test invalid inputs and error handling
- `integration`: Test with FSM controller and token processing
- `performance`: Verify handling of large binary numbers

## References
- `docs/architecture.md`: Rule handler interface and FSM integration
- `audit/golden_test_set.md`: Binary conversion test cases
- Go `strconv` package documentation for number conversion

## Notes for Agent
- Use Go's built-in `strconv.ParseInt` with base 2 for conversion
- Implement consistent error handling pattern with other rule handlers
- Consider edge cases like leading zeros and maximum integer size

## PROMPT — FULL 4-STEP FLOW

```markdown
You are executing **Rule Handler — Binary (TASK-005)** for go-reloaded.

### Step 1 — Analyze & Confirm
- Review FSM controller interface from TASK-003
- Study hexadecimal handler pattern from TASK-004 for consistency
- Examine `audit/golden_test_set.md` for binary conversion test cases
- WAIT for user confirmation before proceeding

### Step 2 — Write the Tests (TDD)
- Create unit tests for valid binary conversions
- Test error handling for invalid binary strings
- Test integration with FSM rule routing
- Prepare edge case tests (empty strings, large numbers)

### Step 3 — Implement the Code
- Build binary conversion handler using Go's strconv package
- Implement error handling and validation
- Integrate with FSM controller interface
- Add logging and debugging support

### Step 4 — Validate & QA
- Run all binary conversion tests
- Test integration with FSM and tokenizer
- Validate error handling and edge cases
- Check performance with large binary inputs
- If verification passes, output: **"✅ Rule Handler — Binary (TASK-005) self-verified. Ready for review."**
```