package recusion

// solveNQueens returns all distinct solutions to the n-queens puzzle.
// Each solution is represented as a slice of strings, where 'Q' denotes a
// queen and '.' denotes an empty space. We use backtracking (recursion) to
// explore valid placements row by row.
func solveNQueens(n int) [][]string {
	var res [][]string

	// initialize empty board
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	// check whether placing a queen at (row,col) is safe
	var isValid func(row, col int) bool
	isValid = func(row, col int) bool {
		// column
		for i := 0; i < row; i++ {
			if board[i][col] == 'Q' {
				return false
			}
		}
		// diagonal left-up
		for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if board[i][j] == 'Q' {
				return false
			}
		}
		// diagonal right-up
		for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
			if board[i][j] == 'Q' {
				return false
			}
		}
		return true
	}

	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			// record current board as a solution
			sol := make([]string, n)
			for i := 0; i < n; i++ {
				sol[i] = string(board[i])
			}
			res = append(res, sol)
			return
		}

		for col := 0; col < n; col++ {
			if isValid(row, col) {
				board[row][col] = 'Q'
				dfs(row + 1)
				board[row][col] = '.'
			}
		}
	}

	dfs(0)
	return res
}
