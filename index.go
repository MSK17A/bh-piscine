package piscine

func Index(s string, toFind string) int {

	for count, word := range s {
		if word == rune(toFind[0]) {
			return count
		}
	}
	return -1
}
