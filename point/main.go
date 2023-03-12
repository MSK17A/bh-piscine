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

/*func StrLen(s string) int {
	runes_counter := 0

	for i := range s {
		if s[i] > 0 {
			runes_counter++
		}
	}
	return runes_counter
}*/

func PrintStr(s string) {
	for _, word := range s {
		z01.PrintRune(word)
	}
}

func PrintNbr(n int) {
	x := '0'
	y := '0'

	for i := 0; i < n/10; i++ {
		x++
	}
	z01.PrintRune(x)
	for i := 0; i < n%10; i++ {
		y++
	}
	z01.PrintRune(y)
}

/*
func conv_to_ASCII(num int, sign int) {
	if num == 0 {
		return
	}
	digit := int((num % 10) * sign)
	_num := int(num / 10)
	conv_to_ASCII(_num, sign)
	z01.PrintRune(48 + rune(digit))
}
*/
