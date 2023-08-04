package nqueen

import (
	"fmt"
)

func solution(board [][]bool, row int, cols []bool, ndiag []bool, rdiag []bool, asf string) {

	if row == len(board) {

		fmt.Println(asf)
		fmt.Println(board)
		return
	}

	for col := 0; col < len(board[0]); col++ {

		if cols[col] == false && ndiag[row+col] == false && rdiag[row-col+len(board)-1] == false {
			board[row][col] = true
			cols[col] = true
			ndiag[row+col] = true
			rdiag[row-col+len(board)-1] = true
			solution(board, row+1, cols, ndiag, rdiag, fmt.Sprintf("%s\n%d - %d", asf, row, col))
			board[row][col] = false
			cols[col] = false
			ndiag[row+col] = false
			rdiag[row-col+len(board)-1] = false
		}
	}
}

func ExecuteNquee() {
	fmt.Println("hello")
	n := 4
	board := make([][]bool, n)
	for i := range board {
		board[i] = make([]bool, n)
	}
	var cols = make([]bool, n)
	var ndiag = make([]bool, 2*n-1)
	var rdiag = make([]bool, 2*n-1)

	solution(board, 0, cols, ndiag, rdiag, "")
}
