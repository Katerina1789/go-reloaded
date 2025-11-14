package internal

import "strconv"

func ApplyBin(result []string) {
	if len(result) > 0 {
		last := result[len(result)-1]
		if val, err := strconv.ParseInt(last, 2, 64); err == nil {
			result[len(result)-1] = strconv.FormatInt(val, 10)
		}
	}
}