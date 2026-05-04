package top150i

func gameOfLife(board [][]int) {

	m := len(board)
	n := len(board[0])
	tmpRes := make([][]int, m)
	for i := range tmpRes {
		tmpRes[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isAlive(board, i, j, m, n) {
				tmpRes[i][j] = 1
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = tmpRes[i][j]
		}
	}

}

func isAlive(board [][]int, i int, j int, m int, n int) bool {
	liveNeighbors := 0
	// 1st element
	if isValidCell(i-1, j-1, m, n) {
		liveNeighbors += board[i-1][j-1]
	}
	// 2nd element
	if isValidCell(i-1, j, m, n) {
		liveNeighbors += board[i-1][j]
	}
	// 3rd  element
	if isValidCell(i-1, j+1, m, n) {
		liveNeighbors += board[i-1][j+1]
	}
	// 4th  element
	if isValidCell(i, j+1, m, n) {
		liveNeighbors += board[i][j+1]
	}
	// 5th  element
	if isValidCell(i+1, j+1, m, n) {
		liveNeighbors += board[i+1][j+1]
	}
	// 6th  element
	if isValidCell(i+1, j, m, n) {
		liveNeighbors += board[i+1][j]
	}
	// 7th  element
	if isValidCell(i+1, j-1, m, n) {
		liveNeighbors += board[i+1][j-1]
	}

	// 8th  element
	if isValidCell(i, j-1, m, n) {
		liveNeighbors += board[i][j-1]
	}

	if board[i][j] == 1 {
		if liveNeighbors >= 2 && liveNeighbors <= 3 {
			return true
		}
	}
	if board[i][j] == 0 {
		if liveNeighbors == 3 {
			return true
		}
	}

	return false
}

func isValidCell(i int, j int, m int, n int) bool {
	if i >= 0 && j >= 0 && i < m && j < n {
		return true
	}
	return false
}
