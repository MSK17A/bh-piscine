package piscine

func FirstRune(s string) rune {
	var result rune
	if len(s) < 1 {
		return 0
	}

	for _, word := range s {
		result = word
		break
	}

	return result
}
