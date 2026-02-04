# Reproducibility Guide

## Overview

This document ensures that any peer can implement the `go-reloaded` project from the provided documentation and tests alone.

## Implementation Checklist

### Prerequisites
- [ ] Go 1.21+ installed
- [ ] Understanding of FSM concepts
- [ ] Familiarity with Go testing package

### Step-by-Step Implementation

#### 1. Project Structure Setup
```
go-reloaded/
├── main.go              # CLI entry point
├── internal/            # Internal package
│   ├── fsm.go          # FSM controller
│   ├── hex.go          # (hex) command
│   ├── bin.go          # (bin) command
│   ├── up.go           # (up) command
│   ├── low.go          # (low) command
│   ├── cap.go          # (cap) command
│   ├── range.go        # Range commands
│   └── punctuation.go  # Post-processing
├── docs/               # Documentation
├── audit/              # Audit framework
└── tasks/              # Task breakdown
```

#### 2. FSM States (from internal/fsm.go)
```go
type State int
const (
    Normal State = iota
    QuoteOpen
)
```

#### 3. Core Functions
```go
// Command functions (return bool for success/failure)
func ApplyHex(result []string) bool
func ApplyBin(result []string) bool
func ApplyUp(result []string, count int)
func ApplyLow(result []string, count int)
func ApplyCap(result []string, count int)
func ApplyRangeRule(result []string, rule string)
func IsRangeRule(word string) bool

// Post-processing functions
func FixPunctuation(text string) string
func FixQuotes(text string) string
func FixArticles(text string) string

// FSM controller
func NewFSM() *FSM
func (fsm *FSM) ProcessWord(word string)
func (fsm *FSM) GetOutput() string
```

#### 4. Implementation Order
1. Create internal package structure
2. Implement individual command files (hex.go, bin.go, up.go, low.go, cap.go)
3. Implement range.go for range transformations  
4. Implement punctuation.go for post-processing (quotes, punctuation, articles)
5. Implement fsm.go controller with tokenization logic
6. Create main.go with CLI handling and audit mode
7. Test with `go run cmd/main.go sample.txt result.txt --audit`

#### 5. Key Implementation Details
- **Tokenization**: Handle range patterns like `(up, 2)` as single tokens
- **FSM Processing**: Apply transformations during word processing
- **Post-processing**: Fix punctuation, quotes, and articles after FSM
- **Audit Mode**: Built-in test suite with 12 comprehensive test cases

### Validation Points

#### After Each Task
- [ ] All unit tests pass
- [ ] Code compiles without warnings
- [ ] Follows Go conventions (gofmt, golint)

#### Final Validation
- [ ] All golden test cases pass
- [ ] Audit mode reports 100% success
- [ ] Performance requirements met (<100ms for 1000 words)
- [ ] Memory usage reasonable (no leaks)

### Expected Behavior Examples

#### Rule Processing
```go
// Input: "hello (up) world"
// FSM calls: internal.ApplyUp(result, 1)
// Output: "HELLO world"

// Input: "1E (hex) files"  
// FSM calls: internal.ApplyHex(result)
// Output: "30 files"
```

#### State Transitions
```go
// Input: "'hello (up) world'"
// States: Normal → QuoteOpen → Normal
// FSM applies transformations inside quotes
// Post-processing fixes quote spacing
// Output: "'HELLO world'"
```

### Common Implementation Pitfalls

1. **Token Boundary Issues**: Ensure punctuation is separate tokens
2. **State Management**: Don't forget to reset states after rule application
3. **Range Handling**: Count backwards correctly for (up, n) rules
4. **Quote Nesting**: Handle multiple quote pairs in same text
5. **Invalid Input**: Gracefully handle invalid hex/binary numbers

### Testing Strategy

#### Unit Tests (per rule)
- Valid inputs with expected outputs
- Invalid inputs (should remain unchanged)
- Edge cases (empty strings, boundaries)

#### Integration Tests
- Full paragraphs from golden_test_set.md
- Performance benchmarks
- Memory profiling

#### Audit Tests
- All 12 test cases from golden_test_set.md
- Automated pass/fail reporting with `--audit` flag
- Performance validation

## Success Criteria

A peer implementation is successful when:
- [ ] All tests in golden_test_set.md pass
- [ ] Audit mode shows 100% success rate
- [ ] Performance meets requirements
- [ ] Code is readable and maintainable
- [ ] No external dependencies used

This guide ensures complete reproducibility of the `go-reloaded` project.