package internal

import "strconv"

func ApplyHex(result []string) {
	if len(result) > 0 {
		last := result[len(result)-1]
		if val, err := strconv.ParseInt(last, 16, 64); err == nil {
			result[len(result)-1] = strconv.FormatInt(val, 10)
		}
	}
}