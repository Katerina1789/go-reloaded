# Rule Handler — Casing (Single)

- **ID**: TASK-006
- **Owner**: Backend Developer
- **Size**: S
- **Confidence**: High
- **Hard Dependencies**: TASK-003
- **Soft Dependencies**: TASK-004, TASK-005
- **Related Architecture**: Text transformation, rule handlers

## Mission Profile
Transform previous word casing using `(up)`, `(low)`, and `(cap)` markers. Implement single-word transformations that modify the immediately preceding word.

## Deliverables
- `internal/up.go`, `internal/low.go`, `internal/cap.go` with casing functions
- `ApplyUp()`, `ApplyLow()`, `ApplyCap()` functions
- Integration with FSM controller for single-word transformations
- Unit tests for each casing rule

## Acceptance Criteria
- ✅ `"hello (up)"` → `"HELLO"`
- ✅ `"WORLD (low)"` → `"world"`
- ✅ `"bridge (cap)"` → `"Bridge"`
- ✅ Handles empty strings and edge cases
- ✅ Preserves non-alphabetic characters

## Verification Plan
- `unit`: Test each casing transformation individually
- `integration`: Test with FSM controller
- `edge`: Test empty strings, numbers, punctuation
- `performance`: Verify efficient string operations

## References
- Go `strings` package for case transformations
- `docs/analysis.md`: Casing rule specifications
- Project requirements for text transformation

## PROMPT — FULL 4-STEP FLOW

```markdown
You are executing **Rule Handler — Casing (Single) (TASK-006)** for go-reloaded.

### Step 1 — Analyze & Confirm
- Review casing transformation requirements for up/low/cap
- Understand single-word transformation logic
- WAIT for user confirmation before proceeding

### Step 2 — Write the Tests (TDD)
- Test uppercase, lowercase, capitalize transformations
- Test edge cases: empty strings, mixed content
- Test preservation of non-alphabetic characters

### Step 3 — Implement the Code
- Build ApplyUp(), ApplyLow(), ApplyCap() functions
- Implement using Go strings package
- Ensure proper handling of edge cases

### Step 4 — Validate & QA
- Run all casing transformation tests
- Verify each rule works independently
- Test integration with FSM
- If verification passes, output: **"✅ Rule Handler — Casing (Single) (TASK-006) self-verified. Ready for review."**
```