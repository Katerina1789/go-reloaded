// Package internal provides post-processing functions for go-reloaded.
// Handles punctuation spacing, quote formatting, and article correction (a/an) after FSM processing.
package internal

import (
	"regexp"  // For pattern matching and replacement
	"strings" // For string manipulation
)

// FixPunctuation normalizes spacing around punctuation marks
func FixPunctuation(text string) string {
	// Remove spaces before punctuation: "word ," → "word,"
	re := regexp.MustCompile(`\s+([.,!?:;])`)
	text = re.ReplaceAllString(text, "$1")

	// Add space after punctuation if missing: "word,text" → "word, text"
	re = regexp.MustCompile(`([.,!?:;])([^\s.,!?:;])`)
	text = re.ReplaceAllString(text, "$1 $2")

	// Remove spaces before grouped punctuation: "word ..." → "word..."
	re = regexp.MustCompile(`\s+(\.{2,}|!{2,}|\?{2,}|[!?]{2,})`)
	text = re.ReplaceAllString(text, "$1")

	return text
}

// FixQuotes removes extra spaces inside quotes: "' text '" → "'text'"
func FixQuotes(text string) string {
	// Match quotes with spaces inside and remove the spaces
	re := regexp.MustCompile(`'\s+([^']*?)\s+'`)
	text = re.ReplaceAllString(text, "'$1'") // Keep content, remove spaces
	return text
}

// FixArticles corrects "a" to "an" before vowels and 'h'
func FixArticles(text string) string {
	// Fix lowercase "a" before vowels/h: "a amazing" → "an amazing"
	re := regexp.MustCompile(`\ba\s+([aeiouAEIOUhH])`)
	text = re.ReplaceAllStringFunc(text, func(match string) string {
		parts := strings.Fields(match) // Split "a word" into ["a", "word"]
		if len(parts) == 2 {
			return "an " + parts[1] // Replace "a" with "an"
		}
		return match // Keep original if parsing fails
	})

	// Fix uppercase "A" before vowels/h: "A amazing" → "An amazing"
	re = regexp.MustCompile(`\bA\s+([aeiouAEIOUhH])`)
	text = re.ReplaceAllStringFunc(text, func(match string) string {
		parts := strings.Fields(match) // Split "A word" into ["A", "word"]
		if len(parts) == 2 {
			return "An " + parts[1] // Replace "A" with "An"
		}
		return match // Keep original if parsing fails
	})

	return text
}
