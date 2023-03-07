package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for counter, word := range os.Args[0] {
		if counter > 1 {
			z01.PrintRune(word)
		}
	}
	z01.PrintRune('\n')
}
