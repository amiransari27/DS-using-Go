package recusion

// solveNQueens returns all distinct solutions to the n-queens puzzle.
// Each solution is represented as a slice of strings, where 'Q' denotes a
// queen and '.' denotes an empty space. We use backtracking (recursion) to
// explore valid placements row by row.
func solveNQueens(n int) [][]string {
	var res [][]string

	//board
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}
	//isvalid
	var isValid func(int, int) bool

	isValid = func(row, col int) bool {

		// row
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

	// dfs
	var dfs func(int)

	dfs = func(row int) {

		if row >= n {
			sol := make([]string, n)
			for i := range board {
				sol[i] = string(board[i])
			}

			res = append(res, sol)
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
