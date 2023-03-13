package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) > 3 {
		return
	}
	if args[1] != "+" && args[1] != "-" && args[1] != "/" && args[1] != "*" && args[1] != "%" {
		return
	}

	a := Atoi(args[0])
	b := Atoi(args[2])
	operator := args[1]

	switch operator {
	case "+":
		{
			fmt.Println(a + b)
		}
	case "-":
		{
			fmt.Println(a - b)
		}
	case "/":
		{
			fmt.Println(a / b)
		}
	case "*":
		{
			fmt.Println(a * b)
		}
	case "%":
		{
			fmt.Println(a % b)
		}
	}
}

func Atoi(s string) int {
	num := 0
	sign := 1
	string_length := len(s)
	index := 0

	if len(s) > 0 {
		if s[0] == 45 {
			sign = -1
			index = 1
		} else if s[0] == 43 {
			sign = 1
			index = 1
		}
	} else {
		return 0
	}

	for i := index; i < string_length; i++ {
		if (s[i] < 48) || (s[i] > 57) {
			return 0
		}
		num = num*10 + (int(s[i]) - 48)
	}
	return num * sign
}
