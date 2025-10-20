Analysis Document 

Problem Description

The go-reloaded tool is a text processor written in Go that automatically applies formatting and correction rules to an input file. These rules include number conversions (e.g., from binary or hexadecimal to decimal), capitalization adjustments, article corrections (a/an), and proper placement of punctuation and quotation marks. The goal is to produce a cleaner, grammatically improved version of the original text.

Rule Breakdown with Examples

(hex) – Converts hexadecimal to decimal
Example: 1E (hex) → 30

(bin) – Converts binary to decimal
Example: 10 (bin) → 2

(up) – Converts the previous word to uppercase
Example: go (up) → GO

(low) – Converts the previous word to lowercase
Example: SHOUTING (low) → shouting

(cap) – Capitalizes the first letter of the previous word
Example: bridge (cap) → Bridge

(up, N) – Converts the previous N words to uppercase
Example: so exciting (up, 2) → SO EXCITING

(low, N) – Converts the previous N words to lowercase
Example: THE WINTER (low, 3) → the winter

(cap, N) – Capitalizes the first letter of the previous N words
Example: foolishness (cap, 6) → Foolishness

Punctuation – Must be attached to the preceding word
Example: boring ,what → boring, what

Grouped punctuation – Preserved as-is
Example: thinking ... you → thinking... you

Quotes – Must wrap the word or phrase tightly
Example: ' awesome ' → 'awesome'

a/an – Corrects the article based on the following word
Example: a amazing → an amazing

Architecture Comparison: FSM vs Pipeline

Pipeline:
- Applies rules sequentially in a linear flow
- Easier to implement and extend
- Best suited for independent transformations

FSM (Finite State Machine):
- Models the text as a sequence of states and transitions
- Allows context-aware decisions and rule interactions
- More complex but powerful for handling overlapping or dependent rules

Chosen Architecture: FSM

I chose FSM because it allows for more precise control over how rules interact, especially in cases where punctuation, capitalization, and article correction overlap. It enables state tracking (e.g., inside quotes, after punctuation) and supports context-sensitive transformations. While more complex to implement, FSM offers greater flexibility and accuracy for this type of text processing.