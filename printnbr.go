package piscine

import "github.com/01-edu/z01"

func conv_to_ASCII(num, sign int) {
	if num == 0 {
		return
	}
	digit := int(num % 10 * sign)
	_num := int(num / 10)
	conv_to_ASCII(_num, sign) // Recursive
	z01.PrintRune(48 + rune(digit))
}
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
