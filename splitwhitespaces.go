package piscine

func SplitWhiteSpaces(s string) []string {
	var strArray []string
	var str_temp []rune

	for _, char := range s {
		if char != ' ' && char != '\t' && char != '\n' {
			str_temp = append(str_temp, char)
		} else if str_temp != nil {
			strArray = append(strArray, string(str_temp))
			str_temp = nil
		}
	}

	strArray = append(strArray, string(str_temp))
	return strArray
}
