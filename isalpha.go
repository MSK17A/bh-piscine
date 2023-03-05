package piscine

func IsAlpha(s string) bool {
	if s == "" {
		return false
	}

	for _, word := range s {
		if (word < 65 || word > 90) && (word < 97 || word > 122) && (word < 48 || word > 57) {
			return false
		}
	}
	return true
}
