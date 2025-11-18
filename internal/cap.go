// Package internal provides capitalization transformation for go-reloaded.
// Capitalizes the first letter of previous word(s) when (cap) or (cap, n) rules are encountered.
package internal

import "strings" // For string case conversion

// ApplyCap capitalizes the first letter of the last 'count' words
// If count exceeds available words, capitalizes all words
func ApplyCap(result []string, count int) {
	start := len(result) - count // Calculate starting position
	if start < 0 {               // Don't go before beginning
		start = 0 // Start from first word if count too large
	}
	// Capitalize each word in range
	for i := start; i < len(result); i++ {
		result[i] = strings.Title(result[i]) // Capitalize first letter
	}
}