package piscine

func BasicAtoi(s string) int {
	num := 0

	for i := 0; i < len(s); i++ {
		num = num*10 + (int(s[i]) - 48)
	}
	return num
}

/*func BasicAtoi(s string) int {
	_s := []byte(s)
	count := 0
	for range _s {
		count += 1
	}
	result := 0
	pow := 1
	for i := count - 1; i >= 0; i-- {
		result += ((int(_s[i]) - 48) * pow)
		pow *= 10
	}
	return result
}*/
