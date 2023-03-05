package piscine

func IsPrintable(s string) bool {
	if s == "" {
		return false
	}

	for _, word := range s {
		if word <= 32 {
			return false
		}
	}
	return true
}
