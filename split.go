package piscine

/*
	func Split(s string, sep rune) []string {
		var strArray []string
		var str_temp []rune

		for _, char := range s {
			if char != sep {
				str_temp = append(str_temp, char)
			} else if str_temp != nil {
				strArray = append(strArray, string(str_temp))
				str_temp = nil
			}
		}

		if str_temp != nil {
			strArray = append(strArray, string(str_temp))
		}
		return strArray
	}
*/
func Split(s string, sep string) []string {
	var strArray []string
	var str_temp []rune
	str_len := len(s)
	sep_len := len(sep)

	for idx, char := range s {
		// Check if not end of the string
		if idx+sep_len < str_len {
			// Search for the seperator string
			if s[idx:idx+sep_len] != sep {
				str_temp = append(str_temp, char)
			} else if str_temp != nil {
				strArray = append(strArray, string(str_temp))
				str_temp = nil
			}
		}
	}

	if str_temp != nil {
		strArray = append(strArray, string(str_temp))
	}
	return strArray
}
