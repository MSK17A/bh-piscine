package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for x := '0'; x <= '9'; x++ {
				for y := j + 1; y <= '9'; y++ {

					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(' ')
					z01.PrintRune(x)
					z01.PrintRune(y)
					if i == '9' && j == '9' && x == '9' && y == '9' {
						break
					}
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
