package piscine

func SplitWhiteSpaces(s string) []string {
	var strArray []string
	var str_temp []rune

	for _, char := range s {
		if char != ' ' && char != 10 && char != 9 {
			str_temp = append(str_temp, char)
		} else {
			strArray = append(strArray, string(str_temp))
			str_temp = nil
		}
	}

	strArray = append(strArray, string(str_temp))
	return strArray
}
