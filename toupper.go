package piscine

func ToUpper(s string) string {
	if s == "" {
		return "0"
	}

	s_upper := []rune(s)
	for counter, word := range s {
		if word >= 97 && word <= 122 {
			s_upper[counter] = word - 32
		} else {
			s_upper[counter] = word
		}
	}
	return string(s_upper)
}
