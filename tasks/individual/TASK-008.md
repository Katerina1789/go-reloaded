# Rule Handler — Punctuation

- **ID**: TASK-008
- **Owner**: Backend Developer
- **Size**: M
- **Confidence**: Medium
- **Hard Dependencies**: TASK-003
- **Soft Dependencies**: None
- **Related Architecture**: Rule handlers, text formatting

## Mission Profile
- Implement punctuation normalization and spacing rules
- Attach punctuation marks to previous words without spaces
- Add appropriate spacing after punctuation marks
- Handle various punctuation types: periods, commas, semicolons, colons, exclamation, question marks

## Deliverables
- `internal/rules/punctuation.go` with punctuation formatting logic
- Punctuation detection and classification system
- Spacing normalization for before/after punctuation
- Comprehensive unit tests for all punctuation scenarios
- Integration with FSM controller for automatic application

## Acceptance Criteria
- ✅ `"hello ,world"` correctly formats to `"hello, world"`
- ✅ `"What ?Really !"` correctly formats to `"What? Really!"`
- ✅ `"Yes ;no :maybe"` correctly formats to `"Yes; no: maybe"`
- ✅ Multiple consecutive punctuation marks are handled properly
- ✅ Punctuation at sentence boundaries works correctly

## Verification Plan
- `unit`: Test each punctuation type with various spacing scenarios
- `multiple`: Test consecutive and mixed punctuation marks
- `integration`: Test with FSM controller and token processing
- `boundary`: Test punctuation at start/end of text

## References
- `docs/architecture.md`: Rule handler interface and token processing
- `audit/golden_test_set.md`: Punctuation formatting test cases
- English punctuation spacing conventions and style guides

## Notes for Agent
- Define punctuation character sets for consistent detection
- Implement spacing rules that work with tokenizer output
- Consider interaction with quote handling for proper formatting

## PROMPT — FULL 4-STEP FLOW

```markdown
You are executing **Rule Handler — Punctuation (TASK-008)** for go-reloaded.

### Step 1 — Analyze & Confirm
- Review tokenizer output from TASK-002 for punctuation handling
- Study FSM controller interface from TASK-003
- Examine `audit/golden_test_set.md` for punctuation formatting test cases
- WAIT for user confirmation before proceeding

### Step 2 — Write the Tests (TDD)
- Create unit tests for each punctuation type and spacing rule
- Test multiple consecutive punctuation scenarios
- Test integration with FSM rule processing
- Prepare edge case tests (start/end boundaries, mixed punctuation)

### Step 3 — Implement the Code
- Build punctuation detection and classification system
- Implement spacing normalization logic
- Integrate with FSM controller for automatic processing
- Add support for various punctuation mark types

### Step 4 — Validate & QA
- Run all punctuation formatting tests
- Test integration with FSM and tokenizer
- Validate edge cases and boundary conditions
- Check consistency with English punctuation conventions
- If verification passes, output: **"✅ Rule Handler — Punctuation (TASK-008) self-verified. Ready for review."**
```