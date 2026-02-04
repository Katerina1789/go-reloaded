# Audit Guide

## Functional Requirements

**Allowed packages:** Only standard Go library

## Audit Mode Requirements

**CLI Integration:** `go-reloaded --audit` or `go-reloaded -a`
**Test Suite:** Comprehensive test cases in `audit/`
**Report Format:** Rule categories with ✅/❌ status, failure details, summary statistics
**Performance:** <5 seconds execution time for full audit suite
**Exit Codes:** 0 (all pass), 1 (any failures)

---

## Test Cases

### Test 1 - Casing Transformations

**Input (`sample.txt`):**
```
If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?
```

**Command:**
```bash
go run cmd/main.go result.txt
```

**Expected Output:**
```
If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?
```

**Result:** Pass [ ] Fail [ ]

---

### Test 2 - Number Conversions

**Input (`sample.txt`):**
```
I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure
```

**Command:**
```bash
go run cmd/main.go result.txt
```

**Expected Output:**
```
I have to pack 5 outfits. Packed 26 just to be sure
```

**Result:** Pass [ ] Fail [ ]

---

### Test 3 - Punctuation Normalization

**Input (`sample.txt`):**
```
Don not be sad ,because sad backwards is das . And das not good
```

**Command:**
```bash
go run cmd/main.go result.txt
```

**Expected Output:**
```
Don not be sad, because sad backwards is das. And das not good
```

**Result:** Pass [ ] Fail [ ]

---

### Test 4 - Complex Rules

**Input (`sample.txt`):**
```
harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '
```

**Command:**
```bash
go run cmd/main.go result.txt
```

**Expected Output:**
```
Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'
```

**Result:** Pass [ ] Fail [ ]