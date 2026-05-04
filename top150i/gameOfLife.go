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
	// Check all 8 neighbors
	directions := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	for _, dir := range directions {
		ni, nj := i+dir[0], j+dir[1]
		if isValidCell(ni, nj, m, n) {
			liveNeighbors += board[ni][nj]
		}
	}

	// A live cell survives with 2-3 neighbors; a dead cell becomes alive with 3 neighbors
	if board[i][j] == 1 {
		return liveNeighbors >= 2 && liveNeighbors <= 3
	}
	return liveNeighbors == 3
}

func isValidCell(i int, j int, m int, n int) bool {
	return i >= 0 && j >= 0 && i < m && j < n
}
