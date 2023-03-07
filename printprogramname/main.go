package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]

	for _, word := range name {
		z01.PrintRune(word)
	}
	z01.PrintRune('\n')
}
