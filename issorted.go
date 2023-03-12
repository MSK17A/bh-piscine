package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) < 1 {
		return true
	}
	sign := f(a[0], a[1])

	for i := a[2]; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) != sign {
			return false
		}
	}
	return true
}
