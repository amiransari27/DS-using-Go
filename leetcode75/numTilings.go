package leetcode75

/*
f(n) = f(n-1) + f(n-2) + 2*f(n-3) + 2*f(n-4) .... + + 2*f(0)
f(n-1) =  f(n-2) + f(n-3) + 2*f(n-4) .... + + 2*f(0)

f(n) - f(n-1) = f(n-1) + f(n-3)
f(n) = 2*f(n-1) + f(n-3)

dp[n] = 2*dp[n-1] + dp[n-3]
*/

func NumTilings(n int) int {
	m := 1000000007
	dp := make([]int, n+1)

	for i, _ := range dp {
		dp[i] = 0
	}

	dp[1], dp[2], dp[3] = 1, 2, 5

	for i := 4; i <= n; i++ {
		dp[i] = (2*dp[i-1]%m + dp[i-3]%m) % m
	}

	return dp[n]
}
