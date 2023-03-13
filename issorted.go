package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) < 1 {
		return true
	}
	sign := f(a[0], a[1])

	for i := 1; i < len(a)-1; i++ {
		// If the sign changed, it means two numbers aren't sorted
		if f(a[i], a[i+1]) != sign {
			return false
		}
	}
	return true
}
