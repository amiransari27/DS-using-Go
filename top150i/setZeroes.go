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
