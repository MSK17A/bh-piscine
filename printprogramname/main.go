package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	for _, word := range os.Args[0] {
		z01.PrintRune(word)
	}
	z01.PrintRune('\n')
}
