package piscine

import "github.com/01-edu/z01"

func main() {
	alphabet := [26]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

	for i := 25; i > -1; i-- {

		z01.PrintRune(alphabet[i])
	}
	z01.PrintRune('\n')

}
