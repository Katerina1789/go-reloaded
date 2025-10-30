# Tokenization Engine

- **ID**: TASK-002
- **Owner**: Backend Developer
- **Size**: M
- **Confidence**: High
- **Hard Dependencies**: TASK-001
- **Soft Dependencies**: None
- **Related Architecture**: FSM foundation, text processing

## Mission Profile
- Build the core tokenization engine that splits text into processable tokens while preserving structure
- Implement intelligent parsing that distinguishes between words, punctuation, whitespace, and rule markers
- Create the foundation for FSM state transitions by providing clean token streams
- Ensure tokenizer preserves original text structure for accurate reconstruction

## Deliverables
- `internal/tokenizer/tokenizer.go` with core tokenization logic
- Token type definitions and structures in `internal/types/token.go`
- Whitespace preservation and reconstruction capabilities
- Unit tests covering various text patterns, edge cases, and rule markers
- Integration with file I/O from TASK-001 for end-to-end text processing

## Acceptance Criteria
- ✅ Tokenizes "Hello, world!" into ["Hello", ",", " ", "world", "!"] preserving all elements
- ✅ Correctly identifies rule markers like "(hex)", "(bin)", "(up, 3)" as single tokens
- ✅ Preserves whitespace patterns including multiple spaces, tabs, and newlines
- ✅ Handles punctuation attachment and separation according to grammar rules
- ✅ Reconstructs original text exactly when tokens are rejoined without transformations

## Verification Plan
- `unit`: Test tokenization of simple sentences, complex punctuation, rule markers
- `integration`: Tokenize and reconstruct sample paragraphs, verify byte-perfect matching
- `edge`: Test empty strings, whitespace-only text, malformed rule markers
- `performance`: Benchmark tokenization speed with large text files

## References
- `docs/architecture.md`: Token types and FSM input requirements
- `audit/golden_test_set.md`: Sample texts for tokenization testing
- Go standard library: `strings`, `unicode` packages for text processing

## Notes for Agent
- Design token structure to support FSM state tracking and rule application
- Consider memory efficiency for large text processing
- Maintain clear separation between tokenization and transformation logic

## PROMPT — FULL 4-STEP FLOW

```markdown
You are executing **Tokenization Engine (TASK-002)** for go-reloaded.

### Step 1 — Analyze & Confirm
- Review `docs/architecture.md` for FSM requirements and token structure needs
- Study `audit/golden_test_set.md` for text patterns that need proper tokenization
- Examine TASK-001 file I/O integration points for text input processing
- WAIT for user confirmation before proceeding

### Step 2 — Write the Tests (TDD)
- Create comprehensive unit tests for various tokenization scenarios
- Test rule marker detection and preservation
- Test whitespace handling and reconstruction accuracy
- Prepare integration tests with sample text files

### Step 3 — Implement the Code
- Build tokenizer with proper token type definitions
- Implement whitespace preservation and text reconstruction
- Create clean interfaces for FSM integration
- Add performance optimizations for large text processing

### Step 4 — Validate & QA
- Run all tokenization tests and verify perfect reconstruction
- Test integration with file I/O from TASK-001
- Benchmark performance with various text sizes
- Validate against golden test set samples
- If verification passes, output: **"✅ Tokenization Engine (TASK-002) self-verified. Ready for review."**
```