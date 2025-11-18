// Package internal provides range rule parsing and handling for go-reloaded.
// Detects and processes complex rules like (up, 2), (low, 3), (cap, 4) that affect multiple words.
package internal

import (
	"regexp" // For pattern matching range rules
	"strconv" // For converting string numbers to integers
)

// IsRangeRule checks if word matches range pattern like "(up, 2)" or "(low, 3)"
func IsRangeRule(word string) bool {
	// Pattern matches: (up|low|cap, number) with optional spaces
	return regexp.MustCompile(`^\((up|low|cap),\s*\d+\)$`).MatchString(word)
}

// ApplyRangeRule parses and applies range transformations like "(up, 2)"
func ApplyRangeRule(result []string, word string) {
	// Extract action and count from pattern like "(up, 2)"
	re := regexp.MustCompile(`\((up|low|cap),\s*(\d+)\)`)
	matches := re.FindStringSubmatch(word) // Get captured groups
	if len(matches) == 3 {                // Should have full match + 2 groups
		action := matches[1]              // "up", "low", or "cap"
		count, _ := strconv.Atoi(matches[2]) // Convert number string to int
		// Apply the appropriate transformation
		switch action {
		case "up":
			ApplyUp(result, count) // Apply uppercase to 'count' words
		case "low":
			ApplyLow(result, count) // Apply lowercase to 'count' words
		case "cap":
			ApplyCap(result, count) // Apply capitalization to 'count' words
		}
	}
}