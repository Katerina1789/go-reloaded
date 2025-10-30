# File I/O Setup

- **ID**: TASK-001
- **Owner**: Backend Developer
- **Size**: S
- **Confidence**: High
- **Hard Dependencies**: None
- **Soft Dependencies**: None
- **Related Architecture**: CLI foundation, file operations

## Mission Profile
- Build the foundational CLI interface that reads input files and writes output files unchanged
- Implement command-line argument parsing using os.Args for input/output file paths
- Establish error handling patterns for file operations that will be used throughout the project
- Create the basic structure that all subsequent transformations will build upon

## Deliverables
- `cmd/go-reloaded/main.go` with CLI argument parsing and usage instructions
- File I/O functions in `internal/fileio/fileio.go` using os.ReadFile and os.WriteFile
- Error handling for missing files, permission issues, and invalid arguments
- Basic logging structure for debugging file operations
- Unit tests covering successful file operations and error conditions

## Acceptance Criteria
- ✅ CLI accepts exactly 2 arguments: input file path and output file path
- ✅ Program reads input file content completely and writes identical content to output file
- ✅ Graceful error handling with helpful messages for file not found, permission denied, invalid arguments
- ✅ Usage message displayed when incorrect number of arguments provided
- ✅ All file operations use Go standard library only (no external dependencies)

## Verification Plan
- `unit`: Test file reading/writing with temporary files, error conditions with invalid paths
- `integration`: Run CLI with sample text files, verify byte-for-byte output matching
- `error`: Test with non-existent files, read-only directories, invalid argument counts
- `performance`: Verify memory usage stays reasonable with large files (>1MB)

## References
- Go standard library: `os`, `fmt`, `log` packages
- `docs/architecture.md`: CLI design patterns and error handling approach
- `audit/audit_guide.md`: File operation validation requirements

## Notes for Agent
- Focus on clean, idiomatic Go code structure that will serve as foundation
- Establish consistent error message formatting for user-friendly output
- Consider file size limitations and memory usage patterns

## PROMPT — FULL 4-STEP FLOW

```markdown
You are executing **File I/O Setup (TASK-001)** for go-reloaded.

### Step 1 — Analyze & Confirm
- Review `docs/architecture.md` for CLI design patterns and project structure requirements
- Examine Go standard library documentation for os.Args, os.ReadFile, os.WriteFile
- Confirm understanding of error handling patterns and file operation requirements
- WAIT for user confirmation before proceeding

### Step 2 — Write the Tests (TDD)
- Create unit tests for file reading/writing operations with temporary test files
- Test error conditions: missing files, invalid paths, permission issues
- Test CLI argument parsing with various input scenarios
- Prepare test data files for integration testing

### Step 3 — Implement the Code
- Build main.go with CLI argument parsing and usage instructions
- Implement file I/O functions with proper error handling
- Create clean, modular code structure following Go conventions
- Add logging for debugging file operations

### Step 4 — Validate & QA
- Run all unit tests and verify they pass
- Test CLI with real files to ensure byte-for-byte output matching
- Verify error handling with various failure scenarios
- Check memory usage with larger test files
- If verification passes, output: **"✅ File I/O Setup (TASK-001) self-verified. Ready for review."**
```