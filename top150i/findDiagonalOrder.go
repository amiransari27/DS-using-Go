package top150i

import "slices"

func findDiagonalOrder(mat [][]int) []int {
	m := len(mat)
	n := len(mat[0])
	dMap := make(map[int][]int)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dMap[i+j] = append(dMap[i+j], mat[i][j])
		}
	}

	res := make([]int, 0)

	flip := true
	for i := 0; i < len(dMap); i++ {
		if flip {
			slices.Reverse(dMap[i])
		}
		res = append(res, dMap[i]...)
		flip = !flip
	}

	return res
}
