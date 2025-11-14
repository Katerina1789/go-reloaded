package internal

import "strings"

func ApplyUp(result []string, count int) {
	start := len(result) - count
	if start < 0 {
		start = 0
	}
	for i := start; i < len(result); i++ {
		result[i] = strings.ToUpper(result[i])
	}
}