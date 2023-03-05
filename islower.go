package piscine

func IsLower(s string) bool {
	if s == "" {
		return false
	}

	for _, word := range s {
		if word <= 97 || word >= 122 {
			return false
		}
	}
	return true
}
