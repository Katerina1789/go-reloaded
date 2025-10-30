# Rule Handler — Article Correction

- **ID**: TASK-010
- **Owner**: Backend Developer
- **Size**: S
- **Confidence**: High
- **Hard Dependencies**: TASK-002 (Tokenization), TASK-003 (FSM Controller)
- **Soft Dependencies**: TASK-008 (Punctuation handling for edge cases)
- **Related Architecture**: Rule handlers, linguistic transformations

## Mission Profile
Implement the article correction rule that automatically replaces "a" with "an" when followed by words starting with vowels (a, e, i, o, u) or the letter 'h'. This rule enhances grammatical correctness in the transformed text and must handle case-insensitive matching while preserving the original case of the article.

The handler must integrate seamlessly with the FSM controller and operate on the token stream, looking ahead to the next word to determine if article replacement is necessary.

## Deliverables
- `internal/rules/article.go` implementing article correction logic with vowel/h detection
- Case-insensitive vowel checking while preserving article case ("A" → "An", "a" → "an")
- Integration with FSM controller to apply rule during token processing
- Unit tests covering all vowel cases, 'h' words, edge cases (punctuation, numbers)
- Documentation in `docs/rules/article.md` explaining the rule behavior and examples
- Test cases added to `test/testdata/article_test.txt` with expected outputs

## Acceptance Criteria
- ✅ "a apple" → "an apple" (lowercase vowel)
- ✅ "A elephant" → "An elephant" (uppercase preservation)
- ✅ "a hour" → "an hour" (h-word handling)
- ✅ "a university" → "a university" (no change, 'u' sounds like 'y')
- ✅ "a honest" → "an honest" (h-word)
- ✅ Handles punctuation: "a, apple" → "an, apple"
- ✅ No false positives: "a cat" remains "a cat"
- ✅ Case preservation: "A" → "An", "a" → "an"

## Verification Plan
- `unit`: Test vowel detection (a,e,i,o,u), h-word handling, case preservation, edge cases
- `integration`: Run with FSM controller on multi-sentence paragraphs containing multiple articles
- `edge`: Test with punctuation between article and word, numbers after articles, empty strings
- `regression`: Ensure no interference with other rules (casing, punctuation, quotes)

## References
- English grammar rules: indefinite articles before vowel sounds
- `docs/architecture.md`: FSM integration patterns for rule handlers
- `audit/audit_guide.md`: Article correction validation checklist
- Go standard library: `strings` package for case handling and character checking

## Notes for Agent
- Focus on simple vowel detection (a,e,i,o,u,h) - don't overcomplicate with phonetic rules
- The spec mentions vowels and 'h', not complex phonetic analysis (e.g., "university" edge case may not be required)
- Ensure the rule only applies to standalone "a" or "A", not within words
- Consider lookahead in token stream to check next word's first character
- Preserve original spacing and punctuation around articles

## PROMPT — FULL 4-STEP FLOW

```markdown
You are executing **Rule Handler — Article Correction (TASK-010)** for go-reloaded.

### Step 1 — Analyze & Confirm
- Review `docs/architecture.md` for FSM integration patterns and rule handler structure
- Examine existing rule handlers (TASK-004 through TASK-009) for consistent patterns
- Confirm understanding of article correction requirements: vowels (a,e,i,o,u) and 'h'
- Identify edge cases: punctuation, case preservation, word boundaries
- WAIT for user confirmation before proceeding

### Step 2 — Write the Tests (TDD)
- Create unit tests for basic vowel cases: "a apple" → "an apple"
- Test h-word handling: "a hour" → "an hour"
- Test case preservation: "A elephant" → "An elephant"
- Test edge cases: punctuation between article and word, non-vowel words unchanged
- Test false positives: "a cat" remains unchanged
- Prepare integration test data with multiple articles in paragraphs

### Step 3 — Implement the Code
- Build article.go with vowel/h detection function
- Implement case-preserving replacement logic
- Integrate with FSM controller for token stream processing
- Add lookahead logic to check next word's first character
- Document rule behavior with examples in docs/rules/article.md

### Step 4 — Validate & QA
- Run all unit tests and verify 100% pass rate
- Test integration with FSM on complex paragraphs
- Verify no interference with other rules (punctuation, quotes, casing)
- Check edge cases: empty strings, single characters, numbers
- Run audit checklist for article correction validation
- If verification passes, output: **"✅ Rule Handler — Article Correction (TASK-010) self-verified. Ready for review."**
```
