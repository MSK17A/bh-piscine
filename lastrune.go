package piscine

func LasttRune(s string) rune {
	var result rune
	if len(s) < 0 {
		return 0
	}

	for _, word := range s {
		result = word
	}

	return result
}
