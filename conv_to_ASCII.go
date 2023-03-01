package piscine

import "github.com/01-edu/z01"

func conv_to_ASCII(num int, sign int) {
	if num == 0 {
		return
	}
	digit := int(num % 10 * sign)
	_num := int(num / 10)
	conv_to_ASCII(_num, sign)
	z01.PrintRune(48 + rune(digit))
}
