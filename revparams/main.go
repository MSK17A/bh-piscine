package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	params := os.Args

	for i := len(params)-1; i > 0; i-- {
		for _, word := range params[i] {
			z01.PrintRune(word)
		}
		z01.PrintRune('\n')
	}
}
