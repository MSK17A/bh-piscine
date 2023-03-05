package piscine

func IsNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, word := range s {
		if word < 48 || word > 57 {
			return false
		}
	}
	return true
}
