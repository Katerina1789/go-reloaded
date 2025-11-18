# Rule Handler — Punctuation

- **ID**: TASK-008
- **Owner**: Backend Developer
- **Size**: S
- **Confidence**: High
- **Hard Dependencies**: TASK-003
- **Soft Dependencies**: TASK-010 (post-processing)
- **Related Architecture**: Post-processing, text formatting

## Mission Profile
Normalize punctuation spacing as part of post-processing pipeline. Attach punctuation to preceding word and add space after, while preserving grouped punctuation like `...` and `!!`.

## Deliverables
- Punctuation normalization logic in `internal/punctuation.go`
- `FixPunctuation()` function for post-processing
- Handling for all punctuation types: `. , ! ? : ;`
- Special handling for grouped punctuation
- Unit tests for punctuation spacing rules

## Acceptance Criteria
- ✅ `"hello ,world"` → `"hello, world"`
- ✅ `"text . More"` → `"text. More"`
- ✅ `"thinking ... you"` → `"thinking... you"` (grouped preserved)
- ✅ `"BAMM !!"` → `"BAMM!!"` (grouped preserved)
- ✅ Handles all punctuation types consistently

## Verification Plan
- `unit`: Test individual punctuation types and spacing
- `integration`: Test grouped punctuation preservation
- `edge`: Test multiple spaces, mixed punctuation
- `performance`: Verify efficient text processing

## References
- `docs/analysis.md`: Punctuation rule specifications
- Go `strings` and `regexp` packages
- Post-processing pipeline requirements

## PROMPT — FULL 4-STEP FLOW

```markdown
You are executing **Rule Handler — Punctuation (TASK-008)** for go-reloaded.

### Step 1 — Analyze & Confirm
- Review punctuation spacing rules and grouped handling
- Understand post-processing pipeline integration
- WAIT for user confirmation before proceeding

### Step 2 — Write the Tests (TDD)
- Test individual punctuation spacing normalization
- Test grouped punctuation preservation
- Test edge cases and mixed punctuation

### Step 3 — Implement the Code
- Build FixPunctuation() function
- Implement spacing rules for all punctuation types
- Handle grouped punctuation special cases

### Step 4 — Validate & QA
- Run all punctuation tests
- Verify grouped punctuation preserved correctly
- Test integration with post-processing pipeline
- If verification passes, output: **"✅ Rule Handler — Punctuation (TASK-008) self-verified. Ready for review."**
```