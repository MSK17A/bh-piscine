package main

import (
	"github.com/01-edu/z01"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		PrintStr("Hello\n")
		return
	}

	for _, val := range args {
		content, err := ioutil.ReadFile(val) // For read access.
		if err != nil {
			PrintStr("ERROR: open ")
			PrintStr(val)
			PrintStr(": no such file or directory\n")
		}

		PrintStr(string(content))
	}
}

func PrintStr(s string) {
	for _, word := range s {
		z01.PrintRune(word)
	}
}
