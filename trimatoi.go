package piscine

func TrimAtoi(s string) int {
	num := 0
	sign := 1
	string_rune := []rune(s)
	string_length := strLen(s)
	index := 0
	first_encounter := '0' // A variable to flag what encounterd first (is it a sign or a number)
	if len(s) > 0 {
	} else {
		return 0
	}
	// if first encouter is the sign then take it, otherwise don't
	for i := index; i <= string_length; i++ {
		/* Below are set of questions to decide wether to update the sign or not */
		// If sign is '-' flag it and update the sign variable
		if string_rune[i] == 45 && first_encounter == '0' {
			sign = -1
			first_encounter = '1'
		}
		// If sign is '+' flag it and update the sign variable
		if string_rune[i] == 43 && first_encounter == '0' {
			sign = 1
			first_encounter = '1'
		}
		// If the rune is a number, flag it
		if (string_rune[i] >= 48) && (string_rune[i] <= 57) && (first_encounter == '0') {
			first_encounter = '1'
		}
		// Check if string is a number
		if (string_rune[i] >= 48) && (string_rune[i] <= 57) {
			num = num*10 + (int(string_rune[i]) - 48)
		}
	}
	return num * sign
}
