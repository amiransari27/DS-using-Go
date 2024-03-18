package leetcode75

func UniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}

	return solve_uniquePaths(dp, m, n, 0, 0)

}

func solve_uniquePaths(dp [][]int, m int, n int, c_m int, c_n int) int {
	if c_m > m-1 || c_n > n-1 {
		return 0
	}
	if c_m == m-1 && c_n == n-1 {
		return 1
	}

	if dp[c_m][c_n] != 0 {
		return dp[c_m][c_n]
	}

	right := solve_uniquePaths(dp, m, n, c_m+1, c_n)
	down := solve_uniquePaths(dp, m, n, c_m, c_n+1)

	dp[c_m][c_n] = right + down
	return dp[c_m][c_n]
}
