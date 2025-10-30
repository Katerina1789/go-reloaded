# Audit Mode Documentation

## Overview

Audit mode provides automated validation of the go-reloaded text transformation tool against a comprehensive test suite. It generates detailed pass/fail reports for quality assurance and regression testing.

## Usage

```bash
# Run full audit suite
go-reloaded --audit
go-reloaded -a

# Exit codes:
# 0 - All tests passed
# 1 - One or more tests failed
```

## Test Suite Structure

- **Location:** `audit/test_cases/`
- **Format:** Input/expected output pairs
- **Coverage:** 20+ test scenarios covering all transformation rules
- **Categories:** Hex, binary, casing, punctuation, quotes, articles, integration

## Report Format

```
=== AUDIT REPORT ===

Rule Category: Hexadecimal Conversion
✅ hex_basic: 1E (hex) → 30
✅ hex_lowercase: ff (hex) → 255
❌ hex_invalid: zz (hex) → zz
   Expected: zz
   Actual: 0

Rule Category: Binary Conversion
✅ bin_basic: 101 (bin) → 5
✅ bin_large: 11111111 (bin) → 255

Summary: 4/5 tests passed (80.0%)
```

## Performance Requirements

- Full audit execution: <5 seconds
- Memory usage: Minimal allocations
- CI/CD integration ready

## Test Case Development

Test cases should cover:
- Valid transformations
- Invalid inputs (edge cases)
- Rule interactions
- Boundary conditions
- Performance scenarios