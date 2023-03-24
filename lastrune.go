package piscine

func LastRune(s string) rune {
	var result rune
	if len(s) < 1 {
		return 0
	}

	for _, word := range s {
		result = word
	}

	return result
}
