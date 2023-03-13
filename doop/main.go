package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
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
			result := a + b
			io.WriteString(os.Stdout, string(result))
		}
	case "-":
		{
			io.WriteString(os.Stdout, "Hello World")
		}
	case "/":
		{
			io.WriteString(os.Stdout, "Hello World")
		}
	case "*":
		{
			io.WriteString(os.Stdout, "Hello World")
		}
	case "%":
		{
			io.WriteString(os.Stdout, "Hello World")
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

func conv_to_ASCII(num int, sign int) {
	if num == 0 {
		return
	}
	digit := int((num % 10) * sign)
	_num := int(num / 10)
	conv_to_ASCII(_num, sign)
	z01.PrintRune(48 + rune(digit))
}

func PrintNbr(n int) {
	sign := 1
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		sign *= -1
		z01.PrintRune('-')
	}
	conv_to_ASCII(n, sign)
}
