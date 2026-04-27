package top150i

func rotateImage(matrix [][]int) {

	m := len(matrix)

	for i := 0; i < m; i++ {
		for j := i; j < m; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < m; i++ {
		l, r := 0, len(matrix[i])-1

		for l < r {
			matrix[i][l], matrix[i][r] = matrix[i][r], matrix[i][l]
			l++
			r--
		}
	}
}
