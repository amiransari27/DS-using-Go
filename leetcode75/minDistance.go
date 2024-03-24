package leetcode75

func MinDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	memo := make([][]int, m+1)
	for i := range memo {
		tmp := make([]int, n+1)
		for j := range tmp {
			tmp[j] = -1
		}
		memo[i] = tmp
	}

	return solve_minDistance(memo, word1, word2, m, n)
}

func solve_minDistance(memo [][]int, word1 string, word2 string, i int, j int) int {

	if i == 0 {
		return j
	}

	if j == 0 {
		return i
	}

	if memo[i][j] != -1 {
		return memo[i][j]
	}

	if word1[i-1] == word2[j-1] {
		return solve_minDistance(memo, word1, word2, i-1, j-1)
	}
	insert := 1 + solve_minDistance(memo, word1, word2, i, j-1)
	delete := 1 + solve_minDistance(memo, word1, word2, i-1, j)
	replace := 1 + solve_minDistance(memo, word1, word2, i-1, j-1)

	memo[i][j] = min(insert, delete, replace)

	return memo[i][j]
}
