// Package internal provides capitalization transformation for go-reloaded.
// Capitalizes the first letter of previous word(s) when (cap) or (cap, n) rules are encountered.
package internal

import "unicode" // For Unicode case conversion

// capitalize capitalizes the first letter of a word
func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// ApplyCap capitalizes the first letter of the last 'count' words
// If count exceeds available words, capitalizes all words
func ApplyCap(result []string, count int) {
	start := len(result) - count // Calculate starting position
	if start < 0 {               // Don't go before beginning
		start = 0 // Start from first word if count too large
	}
	// Capitalize each word in range
	for i := start; i < len(result); i++ {
		result[i] = capitalize(result[i]) // Capitalize first letter
	}
}
