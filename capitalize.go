package piscine

func isCapital(word rune) bool {
	if 'A' <= word && word <= 'Z' {
		return true
	}
	return false
}

func isNumber(word rune) bool {
	if '0' <= word && word <= '9' {
		return true
	}
	return false
}

func isLowerCase(word rune) bool {
	if 'a' <= word && word <= 'z' {
		return true
	}
	return false
}

func isAlphaNum(word rune) bool {
	if isCapital(word) || isNumber(word) || isLowerCase(word) {
		return true
	}
	return false
}
func Capitalize(s string) string {
	// If the previous rune contains non alphanumerical characters CAPITALIZE the current rune
	// Other wise lower the current rune
	// If the next rune is not alphanumerical (don't do anything)
	s_rune := []rune(s)
	capitalize_flag := true

	for idx, word := range s_rune {
		if capitalize_flag {
			if isCapital(word) || isNumber(word) {
				capitalize_flag = false
				continue
			} else if isLowerCase(word) {
				s_rune[idx] -= 32
				capitalize_flag = false
			}
		} else if !capitalize_flag && isCapital(word) {
			s_rune[idx] += 32
		} else if !isAlphaNum(word) {
			capitalize_flag = true
		}
	}

	return string(s_rune)
}
