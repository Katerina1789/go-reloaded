# go-reloaded

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-blue?style=for-the-badge)](https://opensource.org/licenses/MIT)
[![Tests](https://img.shields.io/badge/Tests-12/12-brightgreen?style=for-the-badge)](audit/golden_test_set.md)

## Overview

A text transformation tool implementing Finite State Machine architecture for context-aware rule processing. Built with Go as part of Zone01 Athens curriculum.

## Features

- **Number Conversion**: Hexadecimal and binary to decimal
- **Case Transformation**: Single word and range-based operations  
- **Text Formatting**: Punctuation normalization and quote handling
- **Grammar Correction**: Automatic article correction (a/an)
- **Built-in Testing**: Comprehensive audit mode with 12 test cases

## Installation

```bash
git clone https://github.com/Katerina1789/go-reloaded.git
cd go-reloaded
go mod tidy
```

## Usage

```bash
# Transform text file
go run . input.txt output.txt

# Run test suite
go run . sample.txt result.txt --audit
```

## Examples

```
Input:  "1E (hex) files and 101 (bin) users"
Output: "30 files and 5 users"

Input:  "make it happen (up, 3)"
Output: "MAKE IT HAPPEN"

Input:  "a amazing story"
Output: "an amazing story"

Input:  "hello , world !"
Output: "hello, world!"
```

## Architecture

```
Text Input → Tokenizer → FSM Controller → Rule Handlers → Post-Processing → Output
```

**Components:**
- **FSM Controller**: Manages state transitions (Normal ↔ QuoteOpen)
- **Rule Handlers**: Process transformations (hex, bin, case, range)
- **Post-Processing**: Fixes punctuation, quotes, and articles

## Project Structure

```
go-reloaded/
├── internal/           # Core implementation
│   ├── fsm.go         # FSM controller
│   ├── hex.go         # Hex conversion
│   ├── bin.go         # Binary conversion
│   ├── up.go          # Uppercase
│   ├── low.go         # Lowercase
│   ├── cap.go         # Capitalization
│   ├── range.go       # Range operations
│   └── punctuation.go # Post-processing
├── docs/              # Documentation
├── audit/             # Test cases & validation
├── tasks/             # Task breakdown
├── agents.md          # AI agent instructions
└── main.go            # CLI entry point
```

## Testing

The project includes comprehensive testing with 12 validation scenarios:

```bash
go run . sample.txt result.txt --audit

✅ Hex conversion: PASS
✅ Binary conversion: PASS
✅ Case transformations: PASS
✅ Range operations: PASS
✅ Punctuation normalization: PASS
✅ Quote formatting: PASS
✅ Article correction: PASS

Audit Results: 12/12 tests passed
```

## Documentation

- [Architecture Guide](docs/architecture.md) - FSM design and implementation
- [Requirements Analysis](docs/analysis.md) - Problem analysis and rule specifications
- [Test Cases](audit/golden_test_set.md) - Comprehensive validation scenarios
- [AI Usage](docs/ai_usage.md) - Transparency documentation

## License

MIT License - see [LICENSE](LICENSE) file for details.

---

**Zone01 Athens Project**