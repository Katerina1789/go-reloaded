# Audit Mode

- **ID**: TASK-013
- **Owner**: QA Lead / Backend Developer
- **Size**: M
- **Confidence**: Medium
- **Hard Dependencies**: TASK-011 (Integration tests), TASK-012 (CLI & Error Handling)
- **Soft Dependencies**: All rule handler tasks (TASK-004 through TASK-010)
- **Related Architecture**: Testing framework, validation automation, quality assurance

## Mission Profile
Implement an audit mode that runs the go-reloaded tool against a comprehensive test suite and generates a detailed pass/fail checklist report. This mode enables automated validation of all transformation rules, helps identify regressions, and provides a quality gate for releases.

The audit mode should execute predefined test cases, compare outputs against expected results, and produce a human-readable report showing which rules pass/fail with specific examples of failures for debugging.

## Deliverables
- Audit mode flag `--audit` or `-a` added to CLI in `cmd/go-reloaded/main.go`
- Audit test suite in `audit/test_cases/` directory with input/expected pairs
- Audit runner in `internal/audit/runner.go` that executes tests and compares results
- Checklist report generator in `internal/audit/report.go` producing formatted output
- Comprehensive test cases covering all rules and edge cases (20+ test scenarios)
- Audit report format: rule name, status (✅/❌), test case details, failure examples
- Documentation in `docs/audit.md` explaining audit mode usage and test case structure
- CI/CD integration for automated audit execution on commits

## Acceptance Criteria
- ✅ Running `go-reloaded --audit` executes full audit suite without requiring file arguments
- ✅ Audit report shows pass/fail status for each rule category: hex, binary, casing, punctuation, quotes, articles
- ✅ Failed tests display: test case name, input, expected output, actual output
- ✅ Summary statistics: X/Y tests passed, overall pass rate percentage
- ✅ Audit mode exits with code 0 if all tests pass, code 1 if any test fails
- ✅ Audit test cases cover all acceptance criteria from TASK-004 through TASK-010
- ✅ Report is human-readable and suitable for inclusion in documentation
- ✅ Audit execution time <5 seconds for full suite

## Verification Plan
- `unit`: Test audit runner logic, report generation, test case loading
- `integration`: Run audit mode with known passing/failing test cases, verify report accuracy
- `coverage`: Ensure audit test cases cover all rule handlers and edge cases
- `performance`: Verify audit execution completes in reasonable time (<5s)
- `ci-cd`: Integrate audit mode into CI pipeline and verify automated execution

## References
- `audit/audit_guide.md`: Audit requirements and validation checklist
- `audit/golden_test_set.txt`: Reference test cases for audit suite
- `docs/architecture.md`: Testing framework and quality assurance approach
- Go testing: `testing` package patterns for test execution and reporting

## Notes for Agent
- Audit mode should be self-contained and not require external file arguments
- Test cases should be stored in structured format (e.g., JSON or paired .in/.out files)
- Report should be both human-readable (console) and machine-parseable (for CI/CD)
- Consider adding verbose mode (-v) for detailed output during audit
- Audit test cases should be version-controlled and maintained alongside code
- Failed tests should provide enough context for quick debugging
- Consider adding --audit-report flag to save report to file

## PROMPT — FULL 4-STEP FLOW

```markdown
You are executing **Audit Mode (TASK-013)** for go-reloaded.

### Step 1 — Analyze & Confirm
- Review `audit/audit_guide.md` and `audit/golden_test_set.txt` for audit requirements
- Examine existing integration tests (TASK-011) to understand test case structure
- Plan audit test case organization: directory structure, file naming, format
- Design audit report format: checklist layout, failure details, summary statistics
- Identify all rule categories to test: hex, binary, casing (single/range), punctuation, quotes, articles
- Plan CLI integration: --audit flag, exit codes, output formatting
- WAIT for user confirmation before proceeding

### Step 2 — Create Audit Test Suite
- Design 20+ test cases covering all rules and edge cases:
  - Hex conversion: valid/invalid inputs, edge cases
  - Binary conversion: valid/invalid inputs, edge cases
  - Casing: single word, range, edge cases
  - Punctuation: all types, spacing, edge cases
  - Quotes: single/double, nested, edge cases
  - Articles: vowels, h-words, case preservation
  - Integration: multiple rules combined
- Create test case files in audit/test_cases/ with input/expected pairs
- Document test case structure and naming conventions
- Prepare test case index or manifest for audit runner

### Step 3 — Implement Audit Mode
- Add --audit flag to CLI with argument handling
- Build audit runner that loads test cases and executes transformations
- Implement result comparison logic (byte-for-byte or normalized)
- Create report generator with checklist format:
  - Rule category headers
  - Individual test pass/fail status (✅/❌)
  - Failure details: input, expected, actual
  - Summary statistics and pass rate
- Integrate audit mode into main.go with proper exit codes
- Document audit mode usage in docs/audit.md

### Step 4 — Execute & Validate
- Run audit mode with all test cases passing (baseline)
- Introduce intentional failures to verify failure detection and reporting
- Verify report format is clear and provides actionable debugging info
- Test audit mode performance with full suite (<5s execution time)
- Integrate audit mode into CI/CD pipeline
- Run audit mode as final validation before marking task complete
- If audit passes and report is clear, output: **"✅ Audit Mode (TASK-013) self-verified. All audit tests passing. Ready for review."**
```
