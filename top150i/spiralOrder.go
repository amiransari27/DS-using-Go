package top150i

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])

	res := make([]int, 0, m*n)

	dir := 0

	// 0 - means left to right
	// 1 - means top to bottom
	// 0 - means right to left
	// 0 - means bottom to top

	top, bottom := 0, m-1
	left, right := 0, n-1

	for left <= right && top <= bottom {

		if dir == 0 {
			// left to right
			// constant row

			for i := left; i <= right; i++ {
				res = append(res, matrix[top][i])

			}
			top++

		}

		if dir == 1 {
			// top to bottom
			// constant col

			for i := top; i <= bottom; i++ {
				res = append(res, matrix[i][right])

			}
			right--

		}

		if dir == 2 {
			// right to left
			// constant row

			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])

			}
			bottom--

		}

		if dir == 3 {
			// bottom to top
			// constant col

			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])

			}
			left++

		}
		dir++

		if dir == 4 {
			dir = 0
		}

	}

	return res

}
