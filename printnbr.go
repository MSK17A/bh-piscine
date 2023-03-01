package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	sign := 1
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		sign *= -1
		z01.PrintRune('-')
	}
	conv_to_ASCII(n, sign)
}
