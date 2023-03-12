package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}
	if len(args) < 1 {
		fmt.Println("File name missing")
		return
	}

	file, _ := os.Open("quest8.txt") // For read access.

	buffer := make([]byte, 50)
	file.Read(buffer)

	fmt.Println(string(buffer))
}
