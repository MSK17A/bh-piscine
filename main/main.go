package main

import (
	"fmt"
)

func BasicAtoi(s string) int {
	num := 0
	num = int(s[0]-48)*10 + (int(s[1]) - 48)
	for i := 1; i < (len(s) - 1); i++ {
		num = num*10 + (int(s[i+1]) - 48)
	}
	return num
}

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}
