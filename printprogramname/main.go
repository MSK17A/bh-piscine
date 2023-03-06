package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[0]

	for _, word := range arg {
		z01.PrintRune(word)
	}
	z01.PrintRune('\n')
}
