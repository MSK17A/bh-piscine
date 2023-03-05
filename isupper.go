package piscine

func IsUpper(s string) bool {
	if s == "" {
		return false
	}

	for _, word := range s {
		if word <= 65 || word >= 90 {
			return false
		}
	}
	return true
}
