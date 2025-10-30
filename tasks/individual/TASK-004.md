# Rule Handler — Hexadecimal

- **ID**: TASK-004
- **Owner**: Backend Developer
- **Size**: S
- **Confidence**: High
- **Hard Dependencies**: TASK-003
- **Soft Dependencies**: None
- **Related Architecture**: Rule handlers, number conversion

## Mission Profile
- Implement hexadecimal to decimal conversion rule handler for the "(hex)" transformation
- Create robust parsing that validates hexadecimal format and handles edge cases
- Integrate with FSM controller to replace previous token with decimal equivalent
- Establish error handling patterns for invalid hexadecimal inputs

## Deliverables
- `internal/rules/hex.go` with hexadecimal conversion logic
- Validation functions for hexadecimal format checking
- Error handling for invalid hex strings and edge cases
- Unit tests covering valid conversions, invalid inputs, and boundary conditions
- Integration with FSM controller for rule registration and execution

## Acceptance Criteria
- ✅ Converts "1E (hex)" to "30" correctly
- ✅ Handles uppercase and lowercase hex digits (A-F, a-f)
- ✅ Validates hex format and rejects invalid characters
- ✅ Handles edge cases: empty strings, overflow conditions, malformed input
- ✅ Integrates seamlessly with FSM state transitions

## Verification Plan
- `unit`: Test valid hex conversions (0-9, A-F), invalid inputs, edge cases
- `integration`: Test hex rule within FSM context with surrounding text
- `error`: Verify graceful handling of malformed hex strings
- `audit`: Compare outputs against golden test set hex examples

## References
- `docs/architecture.md`: Rule handler interface and FSM integration
- `audit/golden_test_set.md`: Hexadecimal conversion test cases
- Go standard library: `strconv` package for number conversion

## Notes for Agent
- Use Go's built-in hex parsing capabilities for reliability
- Consider case sensitivity and format validation requirements
- Design error handling to be consistent with other rule handlers

## PROMPT — FULL 4-STEP FLOW

```markdown
You are executing **Rule Handler — Hexadecimal (TASK-004)** for go-reloaded.

### Step 1 — Analyze & Confirm
- Review `docs/architecture.md` for rule handler interface requirements
- Study `audit/golden_test_set.md` for hexadecimal conversion examples
- Examine FSM controller from TASK-003 for integration patterns
- WAIT for user confirmation before proceeding

### Step 2 — Write the Tests (TDD)
- Create unit tests for valid hexadecimal conversions
- Test invalid input handling and error conditions
- Test integration with FSM controller and token replacement
- Prepare test cases matching golden test set examples

### Step 3 — Implement the Code
- Build hexadecimal parser with format validation
- Implement conversion logic using Go standard library
- Create FSM integration interface for rule registration
- Add comprehensive error handling and logging

### Step 4 — Validate & QA
- Run all hex conversion tests and verify accuracy
- Test FSM integration with complete text processing
- Validate against audit golden test cases
- Check error handling with various invalid inputs
- If verification passes, output: **"✅ Rule Handler — Hexadecimal (TASK-004) self-verified. Ready for review."**
```