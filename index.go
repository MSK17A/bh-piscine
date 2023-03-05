package piscine

func strLen(s string) int {
	result := 0
	for count, _ := range s {
		result = count
	}
	return result
}

func Index(s string, toFind string) int {
	if len(s) < 1 || len(toFind) < 1 {
		return 0
	}
	if s[0] == 92 {
		return -1
	}
	var s_rune = []rune(s)
	var f_rune = []rune(toFind)
	s_rune_length := strLen(s)

	for i := 0; i < s_rune_length; i++ {
		if s_rune[i] == f_rune[0] {
			return i
		}
	}
	return -1
}
