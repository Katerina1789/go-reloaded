# CLI & Error Handling

- **ID**: TASK-012
- **Owner**: Backend Developer
- **Size**: S
- **Confidence**: High
- **Hard Dependencies**: TASK-001 (File I/O)
- **Soft Dependencies**: TASK-011 (integration tests)
- **Related Architecture**: User interface, error management

## Mission Profile
Comprehensive error handling with user-friendly messages and robust validation. Create production-ready CLI with clear feedback and proper Unix exit codes for all error scenarios.

## Deliverables
- Enhanced CLI in `main.go` with comprehensive error handling
- User-friendly error messages for all failure scenarios
- Proper Unix exit codes: 0=success, 1=error
- Usage instructions and help text
- Error handling for file operations, invalid arguments

## Acceptance Criteria
- ✅ Clear error messages for missing files, permission issues
- ✅ Usage message when incorrect number of arguments provided
- ✅ Proper exit codes: 0 for success, 1 for errors
- ✅ Graceful handling of all error conditions
- ✅ User-friendly feedback for common mistakes

## Verification Plan
- `unit`: Test all error scenarios and exit codes
- `integration`: Test CLI with various invalid inputs
- `usability`: Verify error messages are clear and helpful
- `edge`: Test permission issues, disk space, etc.

## References
- Unix CLI conventions and exit codes
- Go error handling best practices
- User experience guidelines for CLI tools

## PROMPT — FULL 4-STEP FLOW

```markdown
You are executing **CLI & Error Handling (TASK-012)** for go-reloaded.

### Step 1 — Analyze & Confirm
- Review all possible error scenarios and user mistakes
- Plan error message structure and exit code strategy
- WAIT for user confirmation before proceeding

### Step 2 — Write the Tests (TDD)
- Test all error scenarios: missing files, permissions, invalid args
- Test exit codes and error message clarity
- Test usage instructions and help text

### Step 3 — Implement the Code
- Enhance CLI with comprehensive error handling
- Implement clear, user-friendly error messages
- Add proper exit codes and usage instructions

### Step 4 — Validate & QA
- Run all error handling tests
- Verify error messages are clear and actionable
- Test CLI usability with common mistakes
- If verification passes, output: **"✅ CLI & Error Handling (TASK-012) self-verified. Ready for review."**
```