# Integration Test — Full Paragraph

- **ID**: TASK-011
- **Owner**: QA Lead
- **Size**: M
- **Confidence**: High
- **Hard Dependencies**: TASK-010 (post-processing)
- **Soft Dependencies**: All rule handlers (TASK-004 through TASK-009)
- **Related Architecture**: End-to-end testing, performance validation

## Mission Profile
Build comprehensive integration tests that validate the entire text transformation pipeline. Create golden test framework with test cases of increasing complexity to ensure all rules work together correctly.

## Deliverables
- Integration test framework in `audit/golden_test_set.md`
- 12 comprehensive test cases covering all transformation rules
- Performance benchmarks for processing speed
- End-to-end pipeline validation with byte-for-byte output matching
- Test cases covering rule interactions and edge cases

## Acceptance Criteria
- ✅ Golden test framework with input/expected output pairs
- ✅ 12 test cases covering: hex, binary, casing, punctuation, quotes, articles
- ✅ Complex test case mixing multiple rules in single paragraph
- ✅ Byte-for-byte output matching for all test cases
- ✅ Processing time <100ms for 1000-word paragraphs

## Verification Plan
- `unit`: Test individual test case execution
- `integration`: Run full test suite and verify all pass
- `performance`: Benchmark processing speed with large texts
- `regression`: Ensure no rule interactions cause failures

## References
- `docs/analysis.md`: All transformation rule specifications
- Golden test methodology and best practices
- Performance requirements for text processing

## PROMPT — FULL 4-STEP FLOW

```markdown
You are executing **Integration Test — Full Paragraph (TASK-011)** for go-reloaded.

### Step 1 — Analyze & Confirm
- Review all transformation rules and their interactions
- Plan test case structure and complexity progression
- WAIT for user confirmation before proceeding

### Step 2 — Write the Tests (TDD)
- Create 12 comprehensive test cases
- Design complex paragraph mixing multiple rules
- Set up golden test framework structure

### Step 3 — Implement the Code
- Build test execution framework
- Implement performance benchmarking
- Create test case documentation

### Step 4 — Validate & QA
- Run all integration tests and verify 100% pass rate
- Measure and validate performance benchmarks
- Test rule interactions and edge cases
- If verification passes, output: **"✅ Integration Test — Full Paragraph (TASK-011) self-verified. Ready for review."**
```