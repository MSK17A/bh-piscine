package piscine

func BasicAtoi(s string) int {
	num := 0
	if StrLen(s) <= 1 {
		num = int(s[0] - 48)
	} else {
		num = int(s[0]-48)*10 + (int(s[1]) - 48)
		for i := 1; i < (StrLen(s) - 1); i++ {
			num = num*10 + (int(s[i+1]) - 48)
		}
	}
	return num
}
