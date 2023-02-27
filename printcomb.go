package main

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 0; i < 8; i++ {
		z01.PrintRune(rune(i))
	}

	z01.PrintRune('\n')
}

func main() {
	PrintComb()
}
