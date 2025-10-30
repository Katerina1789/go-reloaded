# FSM Controller

- **ID**: TASK-003
- **Owner**: Backend Developer
- **Size**: L
- **Confidence**: Medium
- **Hard Dependencies**: TASK-001, TASK-002
- **Soft Dependencies**: None
- **Related Architecture**: Core FSM implementation, state management

## Mission Profile
- Implement the Finite State Machine controller that manages text transformation states and rule dispatching
- Create state definitions for Normal, QuoteOpen, RulePending, RangePending, PunctuationPending, ArticleCheck
- Build transition logic that routes tokens to appropriate rule handlers based on current state and context
- Establish the central orchestration system that coordinates all text transformations

## Deliverables
- `internal/fsm/controller.go` with main FSM logic and state management
- `internal/fsm/states.go` defining all FSM states and transition rules
- `internal/fsm/context.go` for maintaining transformation context and history
- State transition logging and debugging capabilities
- Integration interfaces for rule handlers and token processing
- Comprehensive unit tests for state transitions and rule routing

## Acceptance Criteria
- ✅ FSM correctly transitions from Normal to RulePending when "(up)" token detected
- ✅ Maintains quote state tracking with proper QuoteOpen/Normal transitions
- ✅ Routes rule tokens to appropriate handlers while preserving context
- ✅ Handles nested and overlapping rule scenarios without state corruption
- ✅ Provides clear debugging output for state transitions and rule applications

## Verification Plan
- `unit`: Test individual state transitions with mock tokens and rule handlers
- `integration`: Process complete sentences with multiple rule applications
- `state`: Verify state consistency under complex transformation scenarios
- `debug`: Test logging and debugging output for troubleshooting support

## References
- `docs/architecture.md`: FSM design, states, and transition specifications
- `audit/golden_test_set.md`: Complex transformation scenarios for FSM testing
- FSM design patterns and Go state machine implementations

## Notes for Agent
- Design FSM to be extensible for future rule additions
- Implement robust error recovery for malformed rule sequences
- Consider performance implications of state tracking and context management

## PROMPT — FULL 4-STEP FLOW

```markdown
You are executing **FSM Controller (TASK-003)** for go-reloaded.

### Step 1 — Analyze & Confirm
- Review `docs/architecture.md` FSM design section for state definitions and transitions
- Study tokenizer output from TASK-002 to understand FSM input format
- Examine `audit/golden_test_set.md` for complex rule interaction scenarios
- WAIT for user confirmation before proceeding

### Step 2 — Write the Tests (TDD)
- Create unit tests for each state transition scenario
- Test rule routing and handler dispatch logic
- Test context preservation across multiple transformations
- Prepare integration tests with tokenizer from TASK-002

### Step 3 — Implement the Code
- Build FSM controller with clean state management
- Implement state transition logic and rule routing
- Create context tracking for transformation history
- Add debugging and logging capabilities

### Step 4 — Validate & QA
- Run all FSM tests and verify correct state transitions
- Test integration with tokenizer and file I/O
- Validate complex rule scenarios from golden test set
- Check performance and memory usage with large inputs
- If verification passes, output: **"✅ FSM Controller (TASK-003) self-verified. Ready for review."**
```