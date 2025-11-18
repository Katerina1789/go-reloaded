# Rule Handler — Hexadecimal

- **ID**: TASK-004
- **Owner**: Backend Developer
- **Size**: S
- **Confidence**: High
- **Hard Dependencies**: TASK-003
- **Soft Dependencies**: None
- **Related Architecture**: Number conversion, rule handlers

## Mission Profile
Convert previous word from hexadecimal to decimal when `(hex)` marker is encountered. Handle valid hex inputs and gracefully ignore invalid ones.

## Deliverables
- `internal/hex.go` with hexadecimal conversion logic
- `ApplyHex()` function that returns bool for success/failure
- Validation for valid hex characters (0-9, A-F, a-f)
- Integration with FSM controller
- Unit tests for valid/invalid hex inputs

## Acceptance Criteria
- ✅ `"1E (hex)"` → `"30"`
- ✅ `"ff (hex)"` → `"255"`
- ✅ `"ZZ (hex)"` → `"ZZ (hex)"` (invalid, unchanged)
- ✅ Case-insensitive hex parsing
- ✅ Returns bool indicating conversion success

## Verification Plan
- `unit`: Test valid hex conversions, invalid input handling
- `integration`: Test with FSM controller
- `edge`: Test empty strings, non-hex characters
- `performance`: Verify efficient parsing

## References
- Go `strconv` package for hex parsing
- `docs/analysis.md`: Hex conversion rule specification
- Project requirements for number conversion

## PROMPT — FULL 4-STEP FLOW

```markdown
You are executing **Rule Handler — Hexadecimal (TASK-004)** for go-reloaded.

### Step 1 — Analyze & Confirm
- Review hex conversion requirements and edge cases
- Understand integration with FSM controller
- WAIT for user confirmation before proceeding

### Step 2 — Write the Tests (TDD)
- Test valid hex conversions (uppercase, lowercase)
- Test invalid input handling
- Test edge cases and boundary conditions

### Step 3 — Implement the Code
- Build ApplyHex() function with validation
- Implement hex parsing using Go standard library
- Ensure proper error handling

### Step 4 — Validate & QA
- Run all hex conversion tests
- Verify invalid inputs remain unchanged
- Test integration with FSM
- If verification passes, output: **"✅ Rule Handler — Hexadecimal (TASK-004) self-verified. Ready for review."**
```