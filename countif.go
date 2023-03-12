package piscine

func CountIf(f func(string) bool, tab []string) int {
	counter := 0
	for _, val := range tab {
		if f(val) {
			counter++
		}
	}
	return counter
}
