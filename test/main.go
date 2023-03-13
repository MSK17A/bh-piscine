package main

import (
	"fmt"
	"piscine"
)

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}

	result1 := piscine.IsSorted(if_first_is_bigger, a1)
	result2 := piscine.IsSorted(if_first_is_bigger, a2)
	result3 := piscine.IsSorted(if_first_is_bigger, []int{-655468, -181010, 176221, -330973, -680281, -464101, -68422, 454601})

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}

func if_first_is_bigger(first, second int) int {
	if first > second {
		return 1
	} else if first == second {
		return 0
	} else {
		return -1
	}
}
