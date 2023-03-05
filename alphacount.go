package piscine

func AlphaCount(s string) int {
	if len(s) < 0 {
		return 0
	}
	result := 0

	for _, word := range s {
		if (word >= 65 && word <= 90) || (word >= 97 && word <= 122) {
			result++
		}
	}

	return result
}
