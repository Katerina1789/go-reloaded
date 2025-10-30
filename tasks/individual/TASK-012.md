# CLI & Error Handling

- **ID**: TASK-012
- **Owner**: Backend Developer
- **Size**: S
- **Confidence**: High
- **Hard Dependencies**: TASK-001 (File I/O Setup)
- **Soft Dependencies**: TASK-011 (Integration tests for error scenarios)
- **Related Architecture**: CLI interface, user experience, error reporting

## Mission Profile
Enhance the CLI interface with comprehensive error handling, user-friendly messages, and robust validation of command-line arguments. Ensure the tool fails gracefully with helpful guidance when encountering invalid inputs, missing files, or incorrect usage patterns.

This task focuses on production-readiness by providing clear feedback to users, preventing cryptic error messages, and establishing consistent error handling patterns throughout the application.

## Deliverables
- Enhanced `cmd/go-reloaded/main.go` with comprehensive argument validation
- User-friendly usage message with examples and syntax explanation
- Error handling for all failure scenarios: missing files, permission issues, invalid paths, malformed input
- Consistent error message formatting with actionable guidance
- Exit codes following Unix conventions (0=success, 1=error, 2=usage error)
- Error handling tests in `test/cli/error_handling_test.go`
- Documentation in `docs/usage.md` with CLI examples and troubleshooting guide

## Acceptance Criteria
- ✅ Running with no arguments displays usage message with examples
- ✅ Running with 1 argument shows error: "Error: missing output file argument" + usage
- ✅ Running with 3+ arguments shows error: "Error: too many arguments" + usage
- ✅ Non-existent input file shows: "Error: input file 'path' not found"
- ✅ Unreadable input file shows: "Error: cannot read input file 'path': permission denied"
- ✅ Unwritable output path shows: "Error: cannot write to output file 'path': permission denied"
- ✅ All error messages are clear, actionable, and user-friendly (no raw Go panic traces)
- ✅ Exit codes: 0 (success), 1 (processing error), 2 (usage error)

## Verification Plan
- `unit`: Test argument parsing with 0, 1, 2, 3+ arguments; verify error messages
- `integration`: Test with non-existent files, read-only files, invalid paths, directory paths
- `usability`: Verify error messages are clear and provide actionable guidance
- `exit-codes`: Confirm correct exit codes for each error scenario
- `edge`: Test with special characters in paths, very long paths, relative vs absolute paths

## References
- Unix exit code conventions: 0=success, 1=general error, 2=usage error
- Go standard library: `os`, `fmt`, `log`, `path/filepath` packages
- `docs/architecture.md`: Error handling patterns and logging approach
- CLI best practices: clear usage messages, helpful error feedback

## Notes for Agent
- Error messages should be concise but informative - tell users what went wrong AND how to fix it
- Usage message should include: syntax, argument descriptions, and 1-2 examples
- Consider adding a --help flag for explicit usage display
- Avoid exposing internal implementation details in error messages
- Test error handling with real-world scenarios users might encounter
- Ensure error messages are consistent in format and tone

## PROMPT — FULL 4-STEP FLOW

```markdown
You are executing **CLI & Error Handling (TASK-012)** for go-reloaded.

### Step 1 — Analyze & Confirm
- Review `cmd/go-reloaded/main.go` from TASK-001 to understand current CLI structure
- Examine `docs/architecture.md` for error handling patterns and logging approach
- Identify all possible error scenarios: argument validation, file operations, processing errors
- Plan error message format and usage message structure
- Review Unix exit code conventions and CLI best practices
- WAIT for user confirmation before proceeding

### Step 2 — Write Error Handling Tests
- Create tests for argument validation: 0, 1, 2, 3+ arguments
- Test file operation errors: non-existent files, permission denied, invalid paths
- Test edge cases: directories as arguments, special characters, empty files
- Verify error message clarity and actionability
- Test exit codes for each error scenario
- Prepare test fixtures: read-only files, non-existent paths, etc.

### Step 3 — Implement Enhanced CLI
- Add comprehensive argument validation with clear error messages
- Create user-friendly usage message with syntax and examples
- Implement consistent error handling for all file operations
- Add proper exit codes following Unix conventions
- Consider adding --help flag for explicit usage display
- Enhance error messages to be actionable and user-friendly
- Document CLI usage and troubleshooting in docs/usage.md

### Step 4 — Validate & QA
- Run all error handling tests and verify correct behavior
- Test with real-world error scenarios: missing files, permission issues
- Verify error messages are clear and provide actionable guidance
- Confirm exit codes match Unix conventions
- Test edge cases: special characters, long paths, relative paths
- Manually test CLI with various invalid inputs to ensure good UX
- If all tests pass and error handling is robust, output: **"✅ CLI & Error Handling (TASK-012) self-verified. Ready for review."**
```
