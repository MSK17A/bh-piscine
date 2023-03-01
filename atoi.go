package piscine

func Atoi(s string) int {
	num := 0
	sign := 1
	sign_count := 0

	for i := 0; i < len(s); i++ {
		if s[i] == 45 || s[i] == 43 {
			sign_count += 1
		} else if ((s[i] < 48) || (s[i] > 57)) && (s[0] != 45 && s[0] != 43) {
			return 0
		}
		if sign_count > 1 {
			return 0
		}
	}
	if s[0] == 45 {
		sign = -1
		for i := 1; i < len(s); i++ {
			num = num*10 + (int(s[i]) - 48)
		}
	} else if s[0] == 43 {
		sign = 1
		for i := 1; i < len(s); i++ {
			num = num*10 + (int(s[i]) - 48)
		}
	} else {
		for i := 0; i < len(s); i++ {
			num = num*10 + (int(s[i]) - 48)
		}
	}

	return num * sign
}
