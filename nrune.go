package piscine

func NRune(s string, n int) rune {
	var result rune
	if n < 0 {
		return 0
	}
	for counter, word := range s {
		if counter == n {
			result = word
			return result
		}
	}
	return 0
}
