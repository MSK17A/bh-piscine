package piscine

func SortWordArr(a []string) {
	temp := "0"
	min_idx := 0
	for i := 0; i < len(a)-1; i++ {
		min_idx = i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[min_idx] {
				min_idx = j
			}
		}

		if min_idx != i {
			temp = a[i]
			a[i] = a[min_idx]
			a[min_idx] = temp
		}
	}
}
