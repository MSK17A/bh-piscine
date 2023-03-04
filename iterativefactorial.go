package piscine

func IterativeFactorial(nb int) int {
	if nb < 1 || nb > 9223372036854775807 {
		return 0
	}
	result := 1
	for i := 1; i <= nb; i++ {
		result *= i
	}

	return result
}
