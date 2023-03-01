package piscine

func BasicAtoi(s string) int {
	num := 0

	for i := 0; i < len(s); i++ {
		num = num*10 + (int(s[i]) - 48)
	}
	return num
}
