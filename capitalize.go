package piscine

func Capitalize(s string) string {
	// If the previous rune contains non alphanumerical characters CAPITALIZE the current rune
	// Other wise lower the current rune

	s_rune := []rune(s)
	s_length := strLen(s)

	for i := 0; i <= s_length; i++ {
		if i > 0 {
			// if it is not alphanumerical
			if (s_rune[i-1] < 65 || s_rune[i-1] > 90) && (s_rune[i-1] < 97 || s_rune[i-1] > 122) && (s_rune[i-1] < 48 || s_rune[i-1] > 57) {
				// If it is lowercase then UPPERCASE IT
				if s_rune[i] >= 97 && s_rune[i] <= 122 {
					s_rune[i] -= 32
				}
				// If ti is UPPERCASE then lowercase IT
			} else if s_rune[i] >= 65 && s_rune[i] <= 90 {
				s_rune[i] += 32
			}
		}
	}
	return string(s_rune)
}
