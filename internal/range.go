package internal

import (
	"regexp"
	"strconv"
)

func IsRangeRule(word string) bool {
	return regexp.MustCompile(`^\((up|low|cap),\s*\d+\)$`).MatchString(word)
}

func ApplyRangeRule(result []string, word string) {
	re := regexp.MustCompile(`\((up|low|cap),\s*(\d+)\)`)
	matches := re.FindStringSubmatch(word)
	if len(matches) == 3 {
		action := matches[1]
		count, _ := strconv.Atoi(matches[2])
		switch action {
		case "up":
			ApplyUp(result, count)
		case "low":
			ApplyLow(result, count)
		case "cap":
			ApplyCap(result, count)
		}
	}
}