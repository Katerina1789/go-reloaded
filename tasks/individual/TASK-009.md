# Rule Handler — Quotes

- **ID**: TASK-009
- **Owner**: Backend Developer
- **Size**: M
- **Confidence**: Medium
- **Hard Dependencies**: TASK-003 (FSM states)
- **Soft Dependencies**: TASK-010 (post-processing)
- **Related Architecture**: State management, post-processing

## Mission Profile
Format quoted phrases by detecting quote boundaries in FSM and fixing spacing in post-processing. Handle both single-word and multi-word quotes with proper spacing removal.

## Deliverables
- Quote state tracking in FSM controller
- `FixQuotes()` function in `internal/punctuation.go`
- Proper spacing removal around quoted content
- Unit tests for single and multi-word quotes
- Integration with FSM state transitions

## Acceptance Criteria
- ✅ `"' awesome '"` → `"'awesome'"`
- ✅ `"' I am the most well-known '"` → `"'I am the most well-known'"`
- ✅ Quote state tracking: Normal ↔ QuoteOpen
- ✅ Handles nested quotes correctly
- ✅ Preserves quote content transformations

## Verification Plan
- `unit`: Test quote spacing removal, state transitions
- `integration`: Test quotes with transformations inside
- `edge`: Test nested quotes, malformed quotes
- `performance`: Verify efficient quote processing

## References
- `docs/analysis.md`: Quote formatting specifications
- FSM state requirements from TASK-003
- Post-processing pipeline integration

## PROMPT — FULL 4-STEP FLOW

```markdown
You are executing **Rule Handler — Quotes (TASK-009)** for go-reloaded.

### Step 1 — Analyze & Confirm
- Review quote formatting rules and FSM state requirements
- Understand quote boundary detection and spacing removal
- WAIT for user confirmation before proceeding

### Step 2 — Write the Tests (TDD)
- Test quote state transitions in FSM
- Test spacing removal for single and multi-word quotes
- Test quotes with transformations inside

### Step 3 — Implement the Code
- Implement quote state tracking in FSM
- Build FixQuotes() function for post-processing
- Handle nested and malformed quotes

### Step 4 — Validate & QA
- Run all quote formatting tests
- Verify FSM state transitions work correctly
- Test integration with transformations
- If verification passes, output: **"✅ Rule Handler — Quotes (TASK-009) self-verified. Ready for review."**
```