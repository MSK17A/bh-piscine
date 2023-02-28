package piscine

func StrRev(s string) string {
	stringLen := len(s)
	reversed := []rune(s)

	for i := stringLen; i > 0; i-- {
		reversed[stringLen-i] = rune(s[i-1])
	}
	return string(reversed)
}
