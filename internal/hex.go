// Package internal provides hexadecimal to decimal conversion for go-reloaded.
// Converts the previous word from hexadecimal format to decimal when (hex) rule is encountered.
package internal

import "strconv" // For number parsing and formatting

// ApplyHex converts the last word from hexadecimal to decimal
// Returns true if conversion succeeded, false if invalid hex
func ApplyHex(result []string) bool {
	if len(result) > 0 { // Make sure we have words to process
		last := result[len(result)-1]                               // Get the last word
		if val, err := strconv.ParseInt(last, 16, 64); err == nil { // Try parsing as hex (base 16)
			result[len(result)-1] = strconv.FormatInt(val, 10) // Convert to decimal string
			return true                                        // Signal success
		}
	}
	return false // Signal failure (invalid hex or no words)
}
