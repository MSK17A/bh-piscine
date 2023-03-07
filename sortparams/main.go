package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	params := os.Args
	SortedParams := []string(params[1:])
	SortParamTable(SortedParams)

	for i := 0; i < len(SortedParams); i++ {
		for _, word := range SortedParams[i] {
			z01.PrintRune(word)
		}
		z01.PrintRune('\n')
	}
}

func SortParamTable(table []string) {
	temp := "0" // Temporary variable to be used fo swapping
	min_idx := 0
	for table_idx_i := 0; table_idx_i < len(table)-1; table_idx_i++ {
		min_idx = table_idx_i
		for table_idx_j := table_idx_i + 1; table_idx_j < len(table); table_idx_j++ {
			if table[table_idx_j] < table[min_idx] {
				min_idx = table_idx_j
			}
		}

		if min_idx != table_idx_i {
			temp = table[table_idx_i]
			table[table_idx_i] = table[min_idx]
			table[min_idx] = temp
		}
	}
}
