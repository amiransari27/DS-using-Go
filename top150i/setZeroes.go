package top150i

func setZeroes(matrix [][]int) {

	m := len(matrix)
	n := len(matrix[0])

	res := make([][]int, m)
	for i := range res {
		res[i] = append([]int{}, matrix[i]...)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {

				for k := 0; k < n; k++ {
					res[i][k] = 0
				}

				for k := 0; k < m; k++ {
					res[k][j] = 0
				}

			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			matrix[i][j] = res[i][j]
		}
	}

}

func setZeroes2(matrix [][]int) {

	m := len(matrix)
	n := len(matrix[0])

	row := make([]bool, m)
	col := make([]bool, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}

		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if row[i] == true || col[j] == true {
				matrix[i][j] = 0
			}

		}
	}

}

func setZeroes3(matrix [][]int) {

	m := len(matrix)
	n := len(matrix[0])
	firstRowImpacted := false
	firstColImpacted := false

	// check first row is impcated
	for col := 0; col < n; col++ {
		if matrix[0][col] == 0 {
			firstRowImpacted = true
			break
		}
	}

	// check first col is impcated
	for row := 0; row < m; row++ {
		if matrix[row][0] == 0 {
			firstColImpacted = true
			break
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if firstColImpacted {
		for row := 0; row < m; row++ {
			matrix[row][0] = 0
		}
	}

	if firstRowImpacted {
		for col := 0; col < n; col++ {
			matrix[0][col] = 0
		}
	}

}
