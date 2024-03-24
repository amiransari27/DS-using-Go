func minDistance(word1 string, word2 string) int {

	return solve_minDistance(word1, word2, len(word1), len(word2))
}

func solve_minDistance(word1 string, word2 string, i int, j int) int {

	if i == 0 {
		return j
	}

	if j == 0 {
		return i
	}

	if word1[i-1] == word2[j-1] {
		return solve_minDistance(word1, word2, i-1, j-1)
	}
	insert := 1 + solve_minDistance(word1, word2, i, j-1)
	delete := 1 + solve_minDistance(word1, word2, i-1, j)
	replace := 1 + solve_minDistance(word1, word2, i-1, j-1)

	return min(insert, delete, replace)
}
