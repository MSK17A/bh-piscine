package piscine

func FindNextPrime(nb int) int {
	if nb < 0 {
		return 2
	}
	if IsPrime(nb) {
		return nb
	}

	return FindNextPrime(nb + 1)
}
