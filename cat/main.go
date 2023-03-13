package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		reader := io.TeeReader(os.Stdin, os.Stdout)
		ioutil.ReadAll(reader)
		os.Stdin.Close()
		os.Stdout.Close()
	}

	for _, val := range args {
		content, err := ioutil.ReadFile(val) // For read access.
		if err != nil {
			PrintStr("ERROR: open ")
			PrintStr(val)
			PrintStr(": no such file or directory\n")
			os.Exit(1)
		}
		PrintStr(string(content))
	}
}

func PrintStr(s string) {
	for _, word := range s {
		z01.PrintRune(word)
	}
}
