// Package internal provides uppercase transformation for go-reloaded.
// Converts the previous word(s) to UPPERCASE when (up) or (up, n) rules are encountered.
package internal

import "strings" // For string case conversion

// ApplyUp converts the last 'count' words to uppercase
// If count exceeds available words, converts all words
func ApplyUp(result []string, count int) {
	start := len(result) - count // Calculate starting position
	if start < 0 {               // Don't go before beginning
		start = 0 // Start from first word if count too large
	}
	// Convert each word in range to uppercase
	for i := start; i < len(result); i++ {
		result[i] = strings.ToUpper(result[i]) // Convert to UPPERCASE
	}
}