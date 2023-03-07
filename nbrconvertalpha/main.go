package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	params := os.Args
	capitalize := false
	starting_index := 1
	number := 0

	// If no params after the --upper argument, terminate the program
	if len(params) < 2 {
		return
	}
	// If there is --upper flag it, and update the starting_index to start reading the params after the --upper argument
	if params[1] == "--upper" {
		capitalize = true
		starting_index = 2
	}

	var alphaMap [26]rune // A variable to map the alpha to a position in the array

	// If the capital flag is set, then change the whole array to a CAPITAL LETTERS, other wise just map the lowercase
	if capitalize {
		for i := 0; i < 26; i++ {
			alphaMap[i] = rune(i+97) - 32
		}
	} else {
		for i := 0; i < 26; i++ {
			alphaMap[i] = rune(i + 97)
		}
	}

	for _, word := range params[starting_index:] {
		number = string_to_int(word)
		/* Because our mapping start from zero, and the real world mappings should start from 1,
		I just subtracted the given number to make it usable with our mappings array */
		number -= 1
		// If the number is withing our map, you can procceed to print the corresponded alpha, otherwise print space ' '
		if 0 <= number && number < 26 {
			z01.PrintRune(alphaMap[number])
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

// A function that converts string to integer number
func string_to_int(s string) int {
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
