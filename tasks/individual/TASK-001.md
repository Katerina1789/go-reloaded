# File I/O Setup

- **ID**: TASK-001
- **Owner**: Backend Developer
- **Size**: S
- **Confidence**: High
- **Hard Dependencies**: None
- **Soft Dependencies**: None
- **Related Architecture**: CLI foundation, file operations

## Mission Profile
Build the foundational CLI interface that reads input files and writes output files unchanged. Implement command-line argument parsing using os.Args for input/output file paths. Establish error handling patterns for file operations that will be used throughout the project.

## Deliverables
- `main.go` with CLI argument parsing and usage instructions
- File I/O functions using os.ReadFile and os.WriteFile
- Error handling for missing files, permission issues, and invalid arguments
- Basic structure that all subsequent transformations will build upon

## Acceptance Criteria
- ✅ CLI accepts exactly 2 arguments: input file path and output file path
- ✅ Program reads input file content and writes identical content to output file
- ✅ Graceful error handling with helpful messages for file errors
- ✅ Usage message displayed when incorrect number of arguments provided
- ✅ All file operations use Go standard library only

## Verification Plan
- `unit`: Test file reading/writing with temporary files, error conditions
- `integration`: Run CLI with sample text files, verify byte-for-byte output matching
- `error`: Test with non-existent files, invalid argument counts
- `performance`: Verify memory usage stays reasonable with large files

## References
- Go standard library: `os`, `fmt` packages
- `docs/architecture.md`: CLI design patterns
- Project requirements: File I/O specifications

## PROMPT — FULL 4-STEP FLOW

```markdown
You are executing **File I/O Setup (TASK-001)** for go-reloaded.

### Step 1 — Analyze & Confirm
- Review Go standard library documentation for os.Args, os.ReadFile, os.WriteFile
- Confirm understanding of CLI argument parsing requirements
- WAIT for user confirmation before proceeding

### Step 2 — Write the Tests (TDD)
- Create tests for file reading/writing operations
- Test CLI argument parsing scenarios
- Test error conditions: missing files, invalid paths

### Step 3 — Implement the Code
- Build main.go with CLI argument parsing
- Implement file I/O functions with error handling
- Create clean, modular code structure

### Step 4 — Validate & QA
- Run all tests and verify they pass
- Test CLI with real files
- Verify error handling works correctly
- If verification passes, output: **"✅ File I/O Setup (TASK-001) self-verified. Ready for review."**
```