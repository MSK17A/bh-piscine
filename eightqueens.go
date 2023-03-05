package piscine

import "github.com/01-edu/z01"

var board [8][8]int

func isSafe(x, y int) bool {
	// Same row
	for i := y - 1; i >= 0; i-- {
		if board[x][i] == 1 {
			return false
		}
	}

	// Check for diagonal down
	for i, j := x+1, y-1; i < 8 && j >= 0; i, j = i+1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}

	// Check for diagonal up
	for i, j := x-1, y-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}

	return true
}

func placeQueens(col int) {
	if col == 8 {
		// We reached final col to put our final queen
		for i := 0; i < 8; i++ {
			for j := 0; j < 8; j++ {
				if board[j][i] == 1 {

					z01.PrintRune(rune(j + 48 + 1))
				}
			}
		}
		z01.PrintRune('\n')
		return
	}

	for i := 0; i < 8; i++ {
		if isSafe(i, col) {
			board[i][col] = 1
			placeQueens(col + 1)
		}
		// Set to zero to backtrack
		board[i][col] = 0
	}

	return
}

func EightQueens() {
	placeQueens(0)
}
