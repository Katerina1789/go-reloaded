# FSM Controller

- **ID**: TASK-003
- **Owner**: Backend Developer
- **Size**: M
- **Confidence**: Medium
- **Hard Dependencies**: TASK-002
- **Soft Dependencies**: None
- **Related Architecture**: State management, rule dispatching

## Mission Profile
Manage state transitions and rule dispatching using FSM with `Normal` and `QuoteOpen` states. Route transformation rules to appropriate handlers based on current state and input tokens.

## Deliverables
- `internal/fsm.go` with FSM controller implementation
- State definitions: `Normal`, `QuoteOpen`
- Rule routing logic for `(hex)`, `(bin)`, `(up)`, `(low)`, `(cap)`, range rules
- State transition handling for quote boundaries
- Integration with tokenization engine

## Acceptance Criteria
- ✅ `"hello (up)"` routes to uppercase handler
- ✅ Quote state tracking: `Normal` → `QuoteOpen` → `Normal`
- ✅ Rules apply correctly in both states
- ✅ FSM processes token stream sequentially
- ✅ Output reconstruction maintains text structure

## Verification Plan
- `unit`: Test state transitions, rule routing
- `integration`: Test with quote boundaries and nested rules
- `edge`: Test malformed input, invalid states
- `performance`: Verify efficient processing for long texts

## References
- `docs/architecture.md`: FSM design patterns
- `docs/analysis.md`: FSM vs Pipeline rationale
- Go patterns for state machines

## PROMPT — FULL 4-STEP FLOW

```markdown
You are executing **FSM Controller (TASK-003)** for go-reloaded.

### Step 1 — Analyze & Confirm
- Review FSM architecture requirements and state definitions
- Understand rule routing and state transition logic
- WAIT for user confirmation before proceeding

### Step 2 — Write the Tests (TDD)
- Test state transitions between Normal and QuoteOpen
- Test rule routing for each transformation type
- Test quote boundary handling

### Step 3 — Implement the Code
- Build FSM controller with state management
- Implement rule dispatching logic
- Create clean interfaces for rule handlers

### Step 4 — Validate & QA
- Run all FSM tests
- Verify state transitions work correctly
- Test rule routing accuracy
- If verification passes, output: **"✅ FSM Controller (TASK-003) self-verified. Ready for review."**
```