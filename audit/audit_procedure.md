# Go Reloaded Audit Procedure

> **Purpose**: Ensure responsible AI use, clear problem analysis & task decomposition, solid testing, and genuine understanding. Audits are part of learning, not a gate to "catch" you, but you must be able to explain and re-create your work.

## 0) Pairing & Logistics

* **Find your pair** and arrange a date.
* **Minimum time:** **≥ 1 hour per project** (plan for 60–75 min).
* **Artifacts to review:**
  * `docs/analysis.md` (problem restatement, rules with examples, Pipeline vs FSM choice & rationale)
  * `audit/golden_test_set.md` (NL test cases incl. tricky/edge cases & a long mixed-rules paragraph)
  * Optional: `docs/ai_usage.md` (prompt strategy), `agents.md` (agent instructions), `tasks/` (task decomposition)
* **After the audit:** Post notes to the repo (see §5 Deliverables).

## 1) What the audit must cover

### A. AI Usage (responsible & transparent)

* Where was AI used (research, drafting, tests, CI, docs)? Provide **prompt snippets or summaries**.
* **Verification**: What did you *accept*, *modify*, or *reject* from AI? How did you check against specs?
* **Understanding test**: Can you explain AI-generated parts and re-derive them if needed?

### B. Task Breakdown / Analysis

* Clear **problem description** in your own words.
* Explicit **rules** with small examples; ambiguities called out with proposed handling.
* **Architecture**: compare **Pipeline vs FSM** and justify your choice for *this* problem.
* Assumptions, edge cases, and limitations documented.

### C. Tests (Golden Set)

* Coverage of core functional cases from audit examples.
* **≥ 5 original** tricky cases crafted by you.
* One **long paragraph** mixing many rules.
* Each test states **expected output** and links back to the rule(s) it validates.

### D. Other Quality Signals

* Writing: concise, structured, and tied to specs.
* **Reproducibility**: Could a peer implement from your docs + tests?
* Collaboration trail: commits, issues, journal notes.

## 2) Live Understanding & Re‑coding

During the audit, the auditor picks **one SMALL, representative change** (e.g., adjust a rule or extend a test) and asks the author to **walk through the reasoning and re‑create** the result live (pseudo‑code or small edits to analysis/tests no full implementation required in analysis‑only phases). If you cannot explain or re‑create, that's a red flag.

## 3) Suggested 60–75′ Flow

1. **Warm‑up (5′)**: Author 2–3′ overview; align scope.
2. **Deep dive (35–45′)**: Sections A→D; probe "why".
3. **Live re‑coding (10–15′)**: §2 exercise.
4. **Wrap‑up (5′)**: Actions & due date.

## 4) Scoring Rubric (0–2 each, total /10)

* AI usage **disclosed & validated** …… 0–2
* Problem & rules **clarity** ……………… 0–2
* **Architecture rationale** (Pipeline vs FSM) … 0–2
* **Test coverage & originality** ……………… 0–2
* **Reproducibility & organization** ………… 0–2

**Outcome:** ✅ Accept / ✏️ Revise & resubmit / ❌ Incomplete

## 5) Deliverables (post‑audit)

* Create `/audits/YYYY‑MM‑DD‑peer.md` with:
  * Final score + outcome
  * **Top 2 strengths / Top 2 improvements**

## 6) Examples & Inspiration

These show useful repo hygiene (docs, CI, pages) and analysis artifacts:

* Example analysis docs & prompts (from a pilot repo):
  * `docs/analysis.md`, `docs/ai_usage.md`, `agents.md`, `audit/golden_test_set.md`
* Repos with documentation/CI and GitHub Pages deployment patterns.

---

## Quick Reference for This Project

### Automated Testing
```bash
# Run all 12 test cases automatically
go run . sample.txt result.txt --audit
```

### Key Validation Points
✅ **Number Conversion**: hex/bin to decimal  
✅ **Case Transformation**: up/low/cap with ranges  
✅ **Punctuation**: Spacing normalization  
✅ **Quotes**: Proper boundary handling  
✅ **Articles**: a/an correction  
✅ **Integration**: Multiple rules combined