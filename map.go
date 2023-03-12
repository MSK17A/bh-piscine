package piscine

func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, len(a))

	for idx, val := range a {
		result[idx] = f(val)
	}

	return result
}
