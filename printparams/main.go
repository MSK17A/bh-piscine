package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	params := os.Args

	for i:=1; i<len(params); i++ {
		for _, word := range params[i] {
			z01.PrintRune(word)
		}
		z01.PrintRune('\n')
	}
}
