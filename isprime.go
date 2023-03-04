package piscine

func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	}
	if nb == 2 {
		return true
	}

	result := false

	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			result = false
			break
		}
		if i == nb-1 {
			result = true
		}
	}

	return result
}
