# Rule Handler — Binary

- **ID**: TASK-005
- **Owner**: Backend Developer
- **Size**: S
- **Confidence**: High
- **Hard Dependencies**: TASK-003
- **Soft Dependencies**: TASK-004
- **Related Architecture**: Number conversion, rule handlers

## Mission Profile
Convert previous word from binary to decimal when `(bin)` marker is encountered. Handle valid binary inputs (0s and 1s only) and gracefully ignore invalid ones.

## Deliverables
- `internal/bin.go` with binary conversion logic
- `ApplyBin()` function that returns bool for success/failure
- Validation for valid binary characters (0, 1 only)
- Integration with FSM controller
- Unit tests for valid/invalid binary inputs

## Acceptance Criteria
- ✅ `"101 (bin)"` → `"5"`
- ✅ `"1010 (bin)"` → `"10"`
- ✅ `"22 (bin)"` → `"22 (bin)"` (invalid, unchanged)
- ✅ Only accepts 0 and 1 characters
- ✅ Returns bool indicating conversion success

## Verification Plan
- `unit`: Test valid binary conversions, invalid input handling
- `integration`: Test with FSM controller
- `edge`: Test empty strings, non-binary characters
- `performance`: Verify efficient parsing

## References
- Go `strconv` package for binary parsing
- `docs/analysis.md`: Binary conversion rule specification
- Project requirements for number conversion

## PROMPT — FULL 4-STEP FLOW

```markdown
You are executing **Rule Handler — Binary (TASK-005)** for go-reloaded.

### Step 1 — Analyze & Confirm
- Review binary conversion requirements and validation rules
- Understand integration with FSM controller
- WAIT for user confirmation before proceeding

### Step 2 — Write the Tests (TDD)
- Test valid binary conversions
- Test invalid input handling (non-0/1 characters)
- Test edge cases and boundary conditions

### Step 3 — Implement the Code
- Build ApplyBin() function with validation
- Implement binary parsing using Go standard library
- Ensure proper error handling

### Step 4 — Validate & QA
- Run all binary conversion tests
- Verify invalid inputs remain unchanged
- Test integration with FSM
- If verification passes, output: **"✅ Rule Handler — Binary (TASK-005) self-verified. Ready for review."**
```