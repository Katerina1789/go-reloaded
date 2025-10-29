# audit_tasks.md

## Audit Task Checklist — `go-reloaded`

This document contains a structured list of audit tasks for verifying the correctness of the `go-reloaded` text transformation tool. Each task includes the input text, expected output, and a pass/fail checkbox for auditors to mark results.

Auditors should use this checklist in combination with `audit_guide.txt` and `golden_test_set.txt` to ensure consistent and accurate evaluations.

---

## ✅ Audit Instructions

- Run the tool using: `go run . input.txt output.txt`
- Compare the contents of `output.txt` with the expected output
- Mark each test as Pass or Fail
- Use only standard Go packages
- Confirm that all transformation rules are applied correctly

---

### Test 1 — Lowercase + Capitalization + Uppercase

**Input File (`sample.txt`):**
If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?


**Expected Output:**
If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?


**Result:**  
Pass [ ] Fail [ ]

---

### Test 2 — Binary + Hexadecimal Conversion

**Input File (`sample.txt`):**
I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure.


**Expected Output:**
I have to pack 5 outfits. Packed 26 just to be sure.


**Result:**  
Pass [ ] Fail [ ]

---

### Test 3 — Punctuation Normalization

**Input File (`sample.txt`):**
Don not be sad ,because sad backwards is das . And das not good

**Expected Output:**
Don not be sad, because sad backwards is das. And das not good

**Result:**  
Pass [ ] Fail [ ]

---

### Test 4 — Capitalization + Quotes + Article Correction

**Input File (`sample.txt`):**
harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '

**Expected Output:**
Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'

**Result:**  
Pass [ ] Fail [ ]

---

### Test 5 — Comprehensive Rule Interaction

**Input File (`sample.txt`):**
' i (cap) wish it need not HAVE HAPPENED (low, 2) in my time , ' said frodo (cap) , ' so do i (up, 2) ... and so do all who live to see such times . ' but that is not for them (up, 3) to decide ; all we have to decide is what to do with the the time that is given to us (cap, 6) : add 1E (hex) years and 101 (bin) days ,and you will still feel a heavy burden .

**Expected Output:**
'I wish it need not have happened in my time,' said Frodo, 'so DO I... and so do all who live to see such times.' but that is NOT FOR THEM to decide; all we have to decide is what to do with the the Time That Is Given To Us: add 30 years and 5 days, and you will still feel an heavy burden.

**Result:**  
Pass [ ] Fail [ ]

---

## Notes for Auditors

- Use natural language understanding to verify rule accuracy
- Pay close attention to edge cases (e.g., punctuation spacing, multi-word quotes, article correction)
- If a test fails, document the reason and suggest improvements
- Refer to `docs/architecture.md` for rule logic and FSM behavior
- Use `audit_guide.txt` for rule-by-rule verification

---

This checklist ensures consistent and thorough auditing of the `go-reloaded` tool. All contributors should use it during peer review and final evaluation.
