package piscine

func StrLen(s string) int {
	runes_counter := 0

	for i := range s {
		if s[i] > 0 {
			runes_counter++
		}
	}
	return runes_counter
}
