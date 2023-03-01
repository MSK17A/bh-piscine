package piscine

func BasicAtoi(s string) int {
	if s == 0 {
		return
	}
	BasicAtoi(s / 10)
	num = s*10 + s%10
}
