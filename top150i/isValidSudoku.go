package top150i

import "fmt"

func isValidSudoku(board [][]byte) bool {

	// check rows
	for row := 0; row < 9; row++ {
		num_set := make([]bool, 10)
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				continue
			}
			fmt.Printf("Type: %T\n", board[row][col])
			num := int(board[row][col] - '0')
			if num_set[num] {
				return false
			}
			num_set[num] = true
		}
	}

	// check coloums

	for col := 0; col < 9; col++ {
		num_set := make([]bool, 10)
		for row := 0; row < 9; row++ {
			if board[row][col] == '.' {
				continue
			}
			num := int(board[row][col] - '0')

			if num_set[num] {
				return false
			}
			num_set[num] = true
		}

	}

	// check sub grid

	for row := 0; row < 9; row += 3 {
		for col := 0; col < 9; col += 3 {
			if !validataGrid(board, row, row+2, col, col+2) {
				return false
			}
		}
	}

	return true
}

func validataGrid(board [][]byte, sr int, er int, sc int, ec int) bool {

	num_set := make([]bool, 10)
	for r := sr; r <= er; r++ {
		for c := sc; c <= ec; c++ {
			if board[r][c] == '.' {
				continue
			}
			num := int(board[r][c] - '0')

			if num_set[num] {
				return false
			}
			num_set[num] = true

		}
	}

	return true
}
