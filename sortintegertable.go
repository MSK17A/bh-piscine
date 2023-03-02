package piscine

func SortIntegerTable(table []int) {
	temp := 0
	min_idx := 0
	for i := 0; i < len(table)-1; i++ {
		min_idx = i
		for j := i + 1; j < len(table); j++ {
			if table[j] < table[min_idx] {
				min_idx = j
			}
		}

		if min_idx != i {
			temp = table[i]
			table[i] = table[min_idx]
			table[min_idx] = temp
		}
	}
}
