package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}
	if args[1] != "+" && args[1] != "-" && args[1] != "/" && args[1] != "*" && args[1] != "%" {
		return
	}
	for _, char := range args[0] {
		if (char < '0' || char > '9') && char != '-' {
			return
		}
	}
	for _, char := range args[2] {
		if (char < '0' || char > '9') && char != '-' {
			return
		}
	}
	a := Atoi(args[0])
	b := Atoi(args[2])
	operator := args[1]
	switch operator {
	case "+":
		{
			if a+b < 9223372036854775807 || a+b < -9223372036854775807 {
				return
			}
			result := a + b
			PrintNbrStd(result)
			os.Stdout.WriteString("\n")
		}
	case "-":
		{
			if a-b < 9223372036854775807 || a-b < -9223372036854775807 {
				return
			}
			result := a - b
			PrintNbrStd(result)
			os.Stdout.WriteString("\n")
		}
	case "/":
		{
			if b == 0 {
				os.Stdout.WriteString("No division by 0\n")
				return
			}
			result := a / b
			PrintNbrStd(result)
			os.Stdout.WriteString("\n")
		}
	case "*":
		{
			if a*b < 9223372036854775807 || a*b < -9223372036854775807 {
				return
			}
			result := a * b
			PrintNbrStd(result)
			os.Stdout.WriteString("\n")
		}
	case "%":
		{
			if b == 0 {
				os.Stdout.WriteString("No modulo by 0\n")
				return
			}
			result := a % b
			PrintNbrStd(result)
			os.Stdout.WriteString("\n")
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
func conv_to_ASCII(num int, sign int) string {
	if num == 0 {
		return "0"
	}
	digit := int((num % 10) * sign)
	_num := int(num / 10)
	// z01.PrintRune(48 + rune(digit))
	return conv_to_ASCII(_num, sign) + string(48+digit)
}
func PrintNbrStd(n int) {
	sign := 1
	if n == 0 {
		os.Stdout.WriteString("0")
		return
	}
	if n < 0 {
		sign *= -1
		os.Stdout.WriteString("-")
	}
	os.Stdout.WriteString(conv_to_ASCII(n, sign)[1:])
}
func AtoiOverflow(a, b, c int64) bool {
	if a < 0 && c < 0 {
		return a*b+c < 0
	} else if a > 0 && c > 0 {
		return a*b+c > 0
	}
	return true
}
func MultiplyOverflow(a, b, c int64) bool {
	prod := a*b + c
	return (prod/b)-c == a
}
func PlusOverflow(a, b int64) bool {
	if a < 0 && b < 0 {
		return a+b < 0
	} else if a > 0 && b > 0 {
		return a+b > 0
	}
	return true
}
func MinusOverflow(a, b int64) bool {
	if a < 0 && b < 0 {
		if b <= a {
			return a-b >= 0
		}
		return a-b < 0
	} else if a > 0 && b > 0 {
		if a <= b {
			return a-b <= 0
		}
		return a-b > 0
	}
	return true
}
