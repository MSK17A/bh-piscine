package piscine

func strLen(s string) int {
	result := 0
	for count := range s {
		result = count
	}
	return result
}

func Index(s string, toFind string) int {
	if len(s) < 1 || len(toFind) < 1 {
		return 0
	}

	s_rune := []rune(s)
	f_rune := []rune(toFind)
	s_rune_length := strLen(s)
	f_rune_length := strLen(toFind)
	Similarity_count := 0

	// Loop through the length of the input string
	for i := 0; i <= s_rune_length; i++ {
		// If the first rune of the toFind string is the same as the input string
		if s_rune[i] == f_rune[0] {
			// Check if the other letters are the same (loop through all the toFind runes)
			for j := 0; j <= f_rune_length; j++ {
				// Just to make sure we don't loop outside of the input string I make this check
				if i+j <= s_rune_length {
					// If the next
					if s_rune[i+j] == f_rune[j] {
						Similarity_count++
					}
					if Similarity_count >= f_rune_length {
						return i
					}
				}
			}
		}
	}
	return -1
}
