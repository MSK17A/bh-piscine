package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	result := 0

	for i := 0; i <= nb; i++ {
		if i*i == nb {
			result = i
			break
		}
	}

	return result
}
