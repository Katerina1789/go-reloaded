# Integration Test — Full Paragraph

- **ID**: TASK-011
- **Owner**: QA Lead / Backend Developer
- **Size**: M
- **Confidence**: High
- **Hard Dependencies**: TASK-001 through TASK-010 (all rule handlers and core components)
- **Soft Dependencies**: None
- **Related Architecture**: End-to-end pipeline, golden test validation

## Mission Profile
Build comprehensive integration tests that validate the entire text transformation pipeline by processing complete paragraphs containing multiple rules applied simultaneously. Use golden test cases to ensure exact output matching and verify that all rules interact correctly without conflicts.

This task represents the critical validation milestone where all individual components are tested together as a cohesive system, ensuring the FSM correctly orchestrates rule application and produces deterministic, correct output.

## Deliverables
- `test/integration/full_pipeline_test.go` with golden test framework
- Golden test data in `test/testdata/golden/` directory:
  - `input_01.txt` through `input_05.txt` (varied complexity)
  - `expected_01.txt` through `expected_05.txt` (exact expected outputs)
- Test cases covering all rule combinations: hex, binary, casing, punctuation, quotes, articles
- Performance benchmarks for processing speed on large paragraphs (>1000 words)
- Integration test documentation in `docs/testing/integration.md`
- CI/CD integration script for automated golden test execution

## Acceptance Criteria
- ✅ Golden test case 1: Simple paragraph with hex, binary, and casing rules
- ✅ Golden test case 2: Complex paragraph with quotes, punctuation, and articles
- ✅ Golden test case 3: Edge cases with multiple rules on same words
- ✅ Golden test case 4: Large paragraph (500+ words) with all rule types
- ✅ Golden test case 5: Stress test with nested quotes, multiple punctuation, range casing
- ✅ All golden tests produce byte-for-byte identical output to expected files
- ✅ Processing time <100ms for 1000-word paragraphs on standard hardware
- ✅ Zero memory leaks or excessive allocations during processing

## Verification Plan
- `golden`: Run all 5 golden test cases, compare output byte-for-byte with expected files
- `performance`: Benchmark processing speed with varying paragraph sizes (100, 500, 1000, 5000 words)
- `stress`: Test with extreme inputs (10,000+ words, deeply nested rules, edge case combinations)
- `regression`: Re-run all unit tests from TASK-001 through TASK-010 to ensure no breakage
- `memory`: Profile memory usage during large file processing to detect leaks

## References
- `audit/golden_test_set.txt`: Reference test cases provided in project specification
- `docs/architecture.md`: FSM pipeline flow and rule interaction patterns
- `audit/audit_guide.md`: Integration test validation checklist
- Go testing: `testing` package, `testing/quick` for property-based tests
- Golden test pattern: compare output files with expected reference files

## Notes for Agent
- Golden tests are the source of truth - outputs must match exactly (including whitespace)
- Create diverse test cases that exercise all rule combinations and edge cases
- Consider rule interaction scenarios: what happens when multiple rules apply to same token?
- Document any discovered edge cases or unexpected behaviors for future reference
- Performance benchmarks help identify optimization opportunities
- Store golden test files in version control for regression detection

## PROMPT — FULL 4-STEP FLOW

```markdown
You are executing **Integration Test — Full Paragraph (TASK-011)** for go-reloaded.

### Step 1 — Analyze & Confirm
- Review `audit/golden_test_set.txt` for reference test cases and expected behaviors
- Examine `docs/architecture.md` to understand complete pipeline flow from input to output
- Analyze all rule handlers (TASK-004 through TASK-010) to identify interaction scenarios
- Plan 5 golden test cases with increasing complexity covering all rule combinations
- Identify potential edge cases: rule conflicts, order dependencies, boundary conditions
- WAIT for user confirmation before proceeding

### Step 2 — Create Golden Test Data
- Design golden test case 1: Basic rules (hex, binary, simple casing)
- Design golden test case 2: Linguistic rules (quotes, punctuation, articles)
- Design golden test case 3: Complex interactions (range casing with quotes, punctuation with articles)
- Design golden test case 4: Large paragraph (500+ words) with diverse rule usage
- Design golden test case 5: Stress test (nested quotes, multiple consecutive rules, edge cases)
- Create input files and manually verified expected output files
- Document test case purpose and coverage in integration.md

### Step 3 — Implement Integration Tests
- Build full_pipeline_test.go with golden test framework
- Implement file comparison logic for byte-for-byte output validation
- Add performance benchmarks for processing speed measurement
- Create helper functions for test data loading and result comparison
- Integrate with existing test suite and CI/CD pipeline

### Step 4 — Execute & Validate
- Run all 5 golden tests and verify exact output matching
- Execute performance benchmarks and document results
- Run memory profiling to detect leaks or excessive allocations
- Re-run all unit tests (TASK-001 through TASK-010) to ensure no regressions
- Document any discovered issues or edge cases
- Generate test coverage report and ensure >90% coverage
- If all tests pass and performance meets criteria, output: **"✅ Integration Test — Full Paragraph (TASK-011) self-verified. All golden tests passing. Ready for review."**
```
