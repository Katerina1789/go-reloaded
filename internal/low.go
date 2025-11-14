package internal

import "strings" // For string case conversion

// ApplyLow converts the last 'count' words to lowercase
// If count exceeds available words, converts all words
func ApplyLow(result []string, count int) {
	start := len(result) - count // Calculate starting position
	if start < 0 {               // Don't go before beginning
		start = 0 // Start from first word if count too large
	}
	// Convert each word in range to lowercase
	for i := start; i < len(result); i++ {
		result[i] = strings.ToLower(result[i]) // Convert to lowercase
	}
}