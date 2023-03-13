package main

import (
	"github.com/01-edu/z01"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		PrintStr("Hello")
		return
	}

	for _, val := range args {

		file, _ := ioutil.ReadFile(val) // For read access.
		PrintStr(string(file))
	}

}

func PrintStr(s string) {
	for _, word := range s {
		z01.PrintRune(word)
	}
}
