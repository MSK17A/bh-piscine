package piscine

func runeIsAlpha(r rune) bool {
	if (r < 65 || r > 90) && (r < 97 || r > 122) && (r < 48 || r > 57) {
		return false
	}
	return true
}

func Capitalize(s string) string {
	// If the previous rune contains non alphanumerical characters CAPITALIZE the current rune
	// Other wise lower the current rune
	// If the next rune is not alphanumerical (don't do anything)
	s_rune := []rune(s)
	s_length := strLen(s)

	if runeIsAlpha(s_rune[1]) {
		// If it is lowercase then UPPERCASE IT
		if s_rune[0] >= 97 && s_rune[0] <= 122 {
			s_rune[0] -= 32
		}
	}

	for i := 1; i <= s_length; i++ {
		// if it is not alphanumerical (the previous rune)
		if i < s_length {
			if !runeIsAlpha(s_rune[i-1]) && runeIsAlpha(s_rune[i+1]) {
				// If it is lowercase then UPPERCASE IT
				if s_rune[i] >= 97 && s_rune[i] <= 122 {
					s_rune[i] -= 32
				}
			}
		} else if s_rune[i] >= 65 && s_rune[i] <= 90 && runeIsAlpha(s_rune[i-1]) {
			s_rune[i] += 32
		}
	}
	return string(s_rune)
}
