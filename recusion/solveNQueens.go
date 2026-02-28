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

	colMap := make(map[int]bool)
	diaRightUp := make(map[int]bool)
	diaLeftUp := make(map[int]bool)

	//isvalid
	var isValid func(int, int) bool

	isValid = func(row, col int) bool {

		if exist := colMap[col]; exist {
			return false
		}

		if exist := diaRightUp[row+col]; exist {
			return false
		}

		if exist := diaLeftUp[row-col]; exist {
			return false
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

			return
		}

		for col := 0; col < n; col++ {
			if isValid(row, col) {
				colMap[col] = true
				diaRightUp[row+col] = true
				diaLeftUp[row-col] = true
				board[row][col] = 'Q'
				dfs(row + 1)
				board[row][col] = '.'
				colMap[col] = false
				diaRightUp[row+col] = false
				diaLeftUp[row-col] = false
			}
		}

	}

	dfs(0)
	return res
}
