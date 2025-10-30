# Rule Handler — Quotes

- **ID**: TASK-009
- **Owner**: Backend Developer
- **Size**: M
- **Confidence**: Medium
- **Hard Dependencies**: TASK-003
- **Soft Dependencies**: TASK-008 (punctuation interaction)
- **Related Architecture**: Rule handlers, state management

## Mission Profile
- Implement quote formatting and spacing normalization
- Detect quote boundaries and manage quote state in FSM
- Remove extra spaces inside quoted phrases
- Handle both single and double quote scenarios

## Deliverables
- `internal/rules/quotes.go` with quote formatting logic
- Quote state management integration with FSM controller
- Space trimming logic for quoted content
- Comprehensive unit tests for quote scenarios
- Support for nested and adjacent quote handling

## Acceptance Criteria
- ✅ `"' awesome '"` correctly formats to `"'awesome'"`
- ✅ `'" hello world "'"` correctly formats to `"'hello world'"`
- ✅ Multiple words in quotes maintain internal spacing
- ✅ Quote state tracking prevents malformed quote handling
- ✅ Integration with punctuation rules works correctly

## Verification Plan
- `unit`: Test single and multi-word quote formatting
- `state`: Test quote state management and boundary detection
- `integration`: Test with FSM controller and punctuation rules
- `nested`: Test complex quote scenarios and edge cases

## References
- `docs/architecture.md`: FSM state management and rule handler interface
- `audit/golden_test_set.md`: Quote formatting test cases
- Typography conventions for quote spacing and formatting

## Notes for Agent
- Implement quote state tracking in FSM controller integration
- Consider interaction with punctuation rules for proper formatting
- Handle edge cases like unmatched quotes and nested scenarios

## PROMPT — FULL 4-STEP FLOW

```markdown
You are executing **Rule Handler — Quotes (TASK-009)** for go-reloaded.

### Step 1 — Analyze & Confirm
- Review FSM controller state management from TASK-003
- Study punctuation handler from TASK-008 for interaction patterns
- Examine `audit/golden_test_set.md` for quote formatting test cases
- WAIT for user confirmation before proceeding

### Step 2 — Write the Tests (TDD)
- Create unit tests for quote boundary detection and formatting
- Test quote state management and transitions
- Test integration with FSM controller and other rules
- Prepare edge case tests (unmatched quotes, nested scenarios)

### Step 3 — Implement the Code
- Build quote detection and formatting logic
- Implement quote state tracking with FSM integration
- Add space trimming and normalization for quoted content
- Handle interaction with punctuation and other rules

### Step 4 — Validate & QA
- Run all quote formatting tests
- Test integration with FSM and other rule handlers
- Validate state management and edge cases
- Check typography compliance and formatting consistency
- If verification passes, output: **"✅ Rule Handler — Quotes (TASK-009) self-verified. Ready for review."**
```