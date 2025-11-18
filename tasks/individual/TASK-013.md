# Audit Mode

- **ID**: TASK-013
- **Owner**: QA Lead
- **Size**: M
- **Confidence**: High
- **Hard Dependencies**: TASK-011 (integration tests), TASK-012 (CLI)
- **Soft Dependencies**: All rule handlers (TASK-004 through TASK-010)
- **Related Architecture**: Quality assurance, automated validation

## Mission Profile
Automated validation system with comprehensive test suite and detailed reporting. Implement self-contained audit mode with `--audit` flag that runs 12 test cases and generates pass/fail reports.

## Deliverables
- Audit mode flag `--audit` in `main.go`
- Built-in test suite with 12 comprehensive test cases
- Automated test execution and result comparison
- Detailed pass/fail reporting with failure diagnostics
- Performance validation and execution time limits

## Acceptance Criteria
- ✅ `go run . sample.txt result.txt --audit` executes full test suite
- ✅ 12 test cases covering all rules: hex, binary, casing, punctuation, quotes, articles
- ✅ Pass/fail status with ✅/❌ indicators for each test
- ✅ Summary statistics: X/12 tests passed
- ✅ Exit code 0 if all pass, 1 if any fail
- ✅ Execution time <5 seconds for full suite

## Verification Plan
- `unit`: Test audit runner logic and report generation
- `integration`: Run audit mode and verify all tests pass
- `performance`: Verify execution completes within time limit
- `reporting`: Ensure failure details are clear and actionable

## References
- `audit/golden_test_set.md`: Reference test cases
- `audit/audit_procedure.md`: Audit requirements
- Quality assurance best practices

## PROMPT — FULL 4-STEP FLOW

```markdown
You are executing **Audit Mode (TASK-013)** for go-reloaded.

### Step 1 — Analyze & Confirm
- Review audit requirements and test case specifications
- Plan audit execution flow and reporting format
- WAIT for user confirmation before proceeding

### Step 2 — Write the Tests (TDD)
- Create 12 built-in test cases covering all rules
- Design audit execution and comparison logic
- Plan report format and failure diagnostics

### Step 3 — Implement the Code
- Add --audit flag to CLI
- Implement test execution and result comparison
- Build detailed reporting with pass/fail indicators

### Step 4 — Validate & QA
- Run audit mode and verify all tests pass
- Test failure scenarios and error reporting
- Validate performance and execution time
- If verification passes, output: **"✅ Audit Mode (TASK-013) self-verified. All 12 tests passing. Ready for review."**
```