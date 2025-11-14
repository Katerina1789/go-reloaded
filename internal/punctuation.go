package internal

import (
	"regexp"
	"strings"
)

func FixPunctuation(text string) string {
	re := regexp.MustCompile(`\s+([.,!?:;])`)
	text = re.ReplaceAllString(text, "$1")
	re = regexp.MustCompile(`([.,!?:;])([^\s.,!?:;])`)
	text = re.ReplaceAllString(text, "$1 $2")
	re = regexp.MustCompile(`\s+(\.{2,}|!{2,}|\?{2,}|[!?]{2,})`)
	text = re.ReplaceAllString(text, "$1")
	return text
}

func FixQuotes(text string) string {
	re := regexp.MustCompile(`'\s+([^']*?)\s+'`)
	text = re.ReplaceAllString(text, "'$1'")
	return text
}

func FixArticles(text string) string {
	re := regexp.MustCompile(`\ba\s+([aeiouAEIOUhH])`)
	text = re.ReplaceAllStringFunc(text, func(match string) string {
		parts := strings.Fields(match)
		if len(parts) == 2 {
			return "an " + parts[1]
		}
		return match
	})
	re = regexp.MustCompile(`\bA\s+([aeiouAEIOUhH])`)
	text = re.ReplaceAllStringFunc(text, func(match string) string {
		parts := strings.Fields(match)
		if len(parts) == 2 {
			return "An " + parts[1]
		}
		return match
	})
	return text
}