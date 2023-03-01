package piscine

func Atoi(s string) int {
	num := 0
	sign := 1
	sign_count := 0
	string_length := len(s)

	for i := 0; i < string_length; i++ {
		if s[i] == 45 || s[i] == 43 {
			sign_count += 1
		} else if ((s[i] < 48) || (s[i] > 57)) && (s[i] != 45 && s[i] != 43) {
			return 0
		}
		if sign_count > 1 {
			return 0
		}
	}
	/*if s[0] == 45 {
		sign = -1
		for i := 1; i < string_length; i++ {
			num = num*10 + (int(s[i]) - 48)
		}
	} else if s[0] == 43 {
		sign = 1
		for i := 1; i < string_length; i++ {
			num = num*10 + (int(s[i]) - 48)
		}
	} else {
		for i := 0; i < string_length; i++ {
			num = num*10 + (int(s[i]) - 48)
		}
	}*/

	return num * sign
}

/*
func Atoi(s string) int {
	_s := []byte(s)
	count := 0
	for range _s {
		count += 1
	}
	result := 0
	pow := 1
	tmp := 0
	sign := 1
	adj_chars := false
	for i := count - 1; i >= 0; i-- {
		tmp = (int(_s[i]) - 48)
		if tmp >= 0 && tmp <= 9 {
			result += (tmp * pow)
			pow *= 10
		} else if !adj_chars && ((tmp == -5 || tmp == -3) && i == 0) {
			adj_chars = true
			if tmp == -3 {
				sign = -1
			}
		} else {
			return 0
		}
	}
	return result * sign
}
*/
