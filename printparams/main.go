package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, param := range os.Args {
		for _, word := range param {
			z01.PrintRune(word)
		}
		z01.PrintRune('\n')
	}
}
