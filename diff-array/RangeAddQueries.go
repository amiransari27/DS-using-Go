package diffarray

func RangeAddQueries(n int, queries [][]int) [][]int {
	diffMat := make([][]int, n)
	for i := range diffMat {
		diffMat[i] = make([]int, n)
	}

	for _, query := range queries {

		r1, c1, r2, c2 := query[0], query[1], query[2], query[3]

		for i := r1; i <= r2; i++ {
			diffMat[i][c1] += 1

			if c2+1 < n {
				diffMat[i][c2+1] -= 1
			}

		}
	}

	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j < n; j++ {
			diffMat[i][j] += sum
			sum = diffMat[i][j]
		}
	}

	return diffMat
}
