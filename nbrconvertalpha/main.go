package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	var alphaMap [26]rune
	for i:=0; i<26; i++ {
		alphaMap[i] = i+97
	}

	for counter, word := range os.Args[0] {
		if counter > 1 {
			z01.PrintRune(word)
		}
	}
}
