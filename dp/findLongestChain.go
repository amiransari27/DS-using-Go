package dp

import "slices"

func FindLongestChain(pairs [][]int) int {
	n := len(pairs)
	t := make([][]int, n)
	for i := range t {
		t[i] = make([]int, n)
		for j := range t[i] {
			t[i][j] = -1
		}
	}

	slices.SortFunc(pairs, func(a, b []int) int {
		return a[0] - b[0]
	})

	return solveFindLongestChain(pairs, 0, -1, t)
}

func solveFindLongestChain(pairs [][]int, i int, p int, t [][]int) int {

	if i >= len(pairs) {
		return 0
	}

	if p != -1 && t[i][p] != -1 {
		return t[i][p]
	}

	var take int

	if p == -1 || pairs[p][1] < pairs[i][0] {
		take = 1 + solveFindLongestChain(pairs, i+1, i, t)
	}

	not_take := solveFindLongestChain(pairs, i+1, p, t)

	if p != -1 {
		t[i][p] = max(take, not_take)
	}

	return max(take, not_take)

}
