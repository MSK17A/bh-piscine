package main

import (
		"github.com/01-edu/z01"
		"os"
)

func main() {
	arg := os.Args[0]

	for _, word := range arg {
		z01.PrintRune(word)
	}
	z01.PrintRune('\n')
}
