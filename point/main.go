package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	PrintStr("x = ")
	PrintNbr(points.x)
	PrintStr(" ")
	PrintStr("y = ")
	PrintNbr(points.y)
	PrintStr("\n")
}

func PrintStr(s string) {
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
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

func conv_to_ASCII(num int, sign int) {
	if num == 0 {
		return
	}
	digit := int((num % 10) * sign)
	_num := int(num / 10)
	conv_to_ASCII(_num, sign)
	z01.PrintRune(48 + rune(digit))
}
