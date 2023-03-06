package piscine

func nbrToString(num int) string {
	var string_num string
	temp_num := num

	for i := 0; temp_num > 0; i++ {
		string_num += string(temp_num%10 + 48)
		temp_num /= 10
	}
	return string_num
}

func StringSort(str string) string {
	temp := rune('0')
	min_idx := 0
	str_runes := []rune(str)
	str_length := StrLen(str)

	for i := 0; i < str_length-1; i++ {
		min_idx = i
		for j := i + 1; j < str_length; j++ {
			if str_runes[j] < str_runes[min_idx] {
				min_idx = j
			}
		}

		if min_idx != i {
			temp = str_runes[i]
			str_runes[i] = str_runes[min_idx]
			str_runes[min_idx] = temp
		}
	}
	return string(str_runes)
}

func PrintNbrInOrder(n int) {
	PrintStr(StringSort(nbrToString(n)))
}
