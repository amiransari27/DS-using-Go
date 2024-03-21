package leetcode75

import "fmt"

func LongestCommonSubsequence(text1 string, text2 string) int {

	memo := make([][]int, 1001)
	for i := range memo {
		tmp := make([]int, 1001)
		for j := range tmp {
			tmp[j] = -1
		}

		memo[i] = tmp
	}

	return solveLCS(memo, text1, text2, 0, 0)
}

func solveLCS(memo [][]int, text1 string, text2 string, i int, j int) int {
	fmt.Println(i, j)
	if i >= len(text1) || j >= len(text2) {
		return 0
	}
	if memo[i][j] != -1 {
		return memo[i][j]
	}

	if text1[i] == text2[j] {
		memo[i][j] = 1 + solveLCS(memo, text1, text2, i+1, j+1)
		return memo[i][j]
	}

	memo[i][j] = max(
		solveLCS(memo, text1, text2, i, j+1),
		solveLCS(memo, text1, text2, i+1, j),
	)

	return memo[i][j]

}

func LongestCommonSubsequenceOptimized(text1 string, text2 string) int {

	m := len(text1)
	n := len(text2)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range dp {
		for j := range dp[i] {

			if i == 0 || j == 0 {
				dp[i][j] = 0
				continue
			}

			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return dp[m][n]
}
