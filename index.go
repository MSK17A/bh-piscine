package piscine

func Index(s string, toFind string) int {
	if len(s) < 1 || len(toFind) < 1 {
		return 0
	}
	for count, word := range s {
		if word == rune(toFind[0]) {
			return count
		}
	}
	return -1
}
