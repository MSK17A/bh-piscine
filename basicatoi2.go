package piscine

func BasicAtoi2(s string) int {
	num := 0

	for i := 0; i < len(s); i++ {
		if (s[i] < 48) || (s[i] > 57) {
			return 0
		}
	}
	for i := 0; i < len(s); i++ {
		num = num*10 + (int(s[i]) - 48)
	}
	return num
}
