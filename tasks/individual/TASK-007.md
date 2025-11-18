# Rule Handler — Casing (Range)

- **ID**: TASK-007
- **Owner**: Backend Developer
- **Size**: M
- **Confidence**: Medium
- **Hard Dependencies**: TASK-006
- **Soft Dependencies**: TASK-002 (tokenization)
- **Related Architecture**: Range transformations, rule handlers

## Mission Profile
Transform previous `n` words casing using range patterns like `(up, 3)`, `(low, 2)`, `(cap, 4)`. Extend casing logic to handle multiple words counting backwards from the rule marker.

## Deliverables
- `internal/range.go` with range transformation logic
- `ApplyRangeRule()` function that parses and applies range rules
- `IsRangeRule()` function to detect range patterns
- Integration with FSM controller and tokenization
- Unit tests for range transformations and edge cases

## Acceptance Criteria
- ✅ `"this is exciting (up, 3)"` → `"THIS IS EXCITING"`
- ✅ `"THE WINTER (low, 2)"` → `"the winter"`
- ✅ `"foolishness (cap, 6)"` → handles boundary correctly
- ✅ Handles cases where range exceeds available words
- ✅ Proper parsing of range syntax `(rule, number)`

## Verification Plan
- `unit`: Test range parsing, boundary conditions
- `integration`: Test with FSM controller and various ranges
- `edge`: Test ranges exceeding word count, invalid syntax
- `performance`: Verify efficient processing for large ranges

## References
- `docs/analysis.md`: Range rule specifications
- Go `strconv` package for number parsing
- Tokenization requirements from TASK-002

## PROMPT — FULL 4-STEP FLOW

```markdown
You are executing **Rule Handler — Casing (Range) (TASK-007)** for go-reloaded.

### Step 1 — Analyze & Confirm
- Review range rule syntax and counting logic
- Understand backward word counting from rule position
- WAIT for user confirmation before proceeding

### Step 2 — Write the Tests (TDD)
- Test range parsing from `(up, 3)` format
- Test backward word counting and transformation
- Test boundary conditions and invalid ranges

### Step 3 — Implement the Code
- Build ApplyRangeRule() with parsing and transformation
- Implement IsRangeRule() for pattern detection
- Handle edge cases and invalid syntax

### Step 4 — Validate & QA
- Run all range transformation tests
- Verify boundary handling works correctly
- Test integration with FSM and tokenization
- If verification passes, output: **"✅ Rule Handler — Casing (Range) (TASK-007) self-verified. Ready for review."**
```