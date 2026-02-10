package dp

import "slices"

func longestStrChain(words []string) int {

	t := make([][]int, len(words))
	for i := range t {
		t[i] = make([]int, len(words))
		for j := range t[i] {
			t[i][j] = -1
		}
	}

	slices.SortFunc(words, func(a, b string) int {
		return len(a) - len(b)
	})

	return solveLongestStrChain(words, 0, -1, t)
}

func solveLongestStrChain(words []string, i int, pi int, t [][]int) int {

	if i == len(words) {
		return 0
	}

	if pi != -1 && t[i][pi] != -1 {
		return t[i][pi]
	}

	take := 0
	if pi == -1 || isPredecessor(words[pi], words[i]) {
		take = 1 + solveLongestStrChain(words, i+1, i, t)
	}
	not_take := solveLongestStrChain(words, i+1, pi, t)

	if pi != -1 {
		t[i][pi] = max(take, not_take)
	}

	return max(take, not_take)
}

func isPredecessor(a string, b string) bool {

	n := len(a)
	m := len(b)
	if m-n != 1 {
		return false
	}

	i, j := 0, 0

	for i < n && j < m {
		if a[i] == b[j] {
			i++
			j++
		} else {
			j++
		}
	}

	return i == n
}
