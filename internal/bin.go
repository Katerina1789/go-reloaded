// Package internal provides binary to decimal conversion for go-reloaded.
// Converts the previous word from binary format to decimal when (bin) rule is encountered.
package internal

import "strconv" // For number parsing and formatting

// ApplyBin converts the last word from binary to decimal
// Returns true if conversion succeeded, false if invalid binary
func ApplyBin(result []string) bool {
	if len(result) > 0 { // Make sure we have words to process
		last := result[len(result)-1]                       // Get the last word
		if val, err := strconv.ParseInt(last, 2, 64); err == nil { // Try parsing as binary (base 2)
			result[len(result)-1] = strconv.FormatInt(val, 10) // Convert to decimal string
			return true                                       // Signal success
		}
	}
	return false // Signal failure (invalid binary or no words)
}