# Analysis Document

This document outlines the problem, transformation rules, and architectural decisions for the `go-reloaded` project.

## Problem Description

The `go-reloaded` tool is a text processor written in Go programming language that automatically applies formatting and correction rules to an input file. These rules include number conversions (e.g., from binary or hexadecimal to decimal), capitalization adjustments, article corrections (a/an) and proper placement of punctuation and quotation marks. The goal is to produce a cleaner, grammatically improved version of the original text or make personalized adjustments.

## Rule Breakdown with Examples
 
| **Rule Name**         | **Description**                                              | **Example**                             |
|-----------------------|--------------------------------------------------------------|-----------------------------------------|
| (hex)                 | Converts hexadecimal to decimal                              | `1E (hex)` → `30`                       |
| (bin)                 | Converts binary to decimal                                   | `10 (bin)` → `2`                        |
| (up)                  | Converts the previous word to uppercase                      | `go (up)` → `GO`                        |
| (low)                 | Converts the previous word to lowercase                      | `SHOUTING (low)` → `shouting`           |
| (cap)                 | Capitalizes the first letter of the previous word            | `bridge (cap)` → `Bridge`               |
| (up, n)               | Converts the previous n words to uppercase                   | `so exciting (up, 2)` → `SO EXCITING`   |
| (low, n)              | Converts the previous n words to lowercase                   | `THE WINTER (low, 3)` → `the winter`    |
| (cap, n)              | Capitalizes the first letter of the previous n words         | `foolishness (cap, 6)` → `Foolishness`  |
| Punctuation           | Must be attached to the preceding word                       | `boring ,what` → `boring, what`         |
| Grouped punctuation   | Preserved as-is                                              | `thinking ... you` → `thinking... you`  |
| Quotes                | Must wrap the word or phrase tightly                         | `' awesome '` → `'awesome'`             |
| a/an                  | Corrects the article based on the following word             | `a amazing` → `an amazing`              |

## Architecture Comparison: Pipeline vs FSM

### Pipeline

- A pipeline applies rules in a fixed, linear sequence.
- Each rule operates independently without considering surrounding context.
- It is simple to implement and works well for isolated transformations like `(hex)` or `(up)`.
- Overlapping rules can conflict, making coordination difficult in complex cases.
- It does not track state, so it cannot detect whether the text is inside quotes or follows punctuation.

### FSM (Finite State Machine)

- An FSM tracks the current state of the text, such as being inside quotes or after punctuation.
- It applies rules dynamically based on contextual information.
- It handles overlapping and grammar-sensitive rules with greater accuracy.
- Building an FSM requires defining states and transitions, making it more complex.
- It offers precise control over rule interactions and is ideal for nuanced text processing.

## Architecture Decision: FSM vs Pipeline

### Pipeline Analysis
**Pros:**
- Simple linear processing
- Easy to implement and debug
- Clear rule separation

**Cons:**
- No context awareness
- Cannot handle overlapping rules (quotes + punctuation)
- Risk of rule conflicts
- Cannot track state for range transformations

### FSM Analysis
**Pros:**
- Context-aware processing
- Handles complex rule interactions
- State tracking for quotes, punctuation, ranges
- Precise control over transformations

**Cons:**
- More complex implementation
- Requires state management

### Final Choice: FSM

**Rationale:** FSM is essential for this project because:
1. **Quote state tracking** - Must know when inside `'text'` to apply rules correctly
2. **Range transformations** - `(up, 3)` must count backwards while respecting quote boundaries
3. **Nested rule interactions** - Rules inside quotes behave differently than outside
4. **Context-sensitive punctuation** - Same punctuation has different spacing rules based on context

**Concrete example justifying FSM:**
Input: `'hello (up, 2) world' and more text ,with punctuation`

**FSM handles:**
- State: Normal → QuoteOpen → Normal
- Applies `(up, 2)` only to words inside quotes: `'HELLO WORLD'`
- Punctuation spacing differs inside vs outside quotes

**Pipeline problems:**
- Cannot track quote boundaries across multiple rules
- Range counting would cross quote boundaries incorrectly
- No way to apply different punctuation rules based on context

This level of contextual coordination requires state management that only FSM provides.