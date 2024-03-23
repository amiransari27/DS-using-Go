package leetcode75

func maxProfit(prices []int, fee int) int {

	dp := make([][]int, len(prices)+1)
    for i := range dp {
        dp[i] = make([]int, 2)
        dp[i][0] = -1
        dp[i][1] = -1
    }

	return solve_maxProfit(dp, prices, fee, 0, 1)
}

func solve_maxProfit(dp [][]int, prices []int, fee int, i int, buy int) int {

	if i >= len(prices) {
		return 0
	}

	profit := 0

	if dp[i][buy] != -1 {
		return dp[i][buy]
	}

	if buy == 1 {
		// make purchase
		cosider := solve_maxProfit(dp, prices, fee, i+1, 0) - prices[i]
		not_consider := solve_maxProfit(dp, prices, fee, i+1, 1)
		profit = max(cosider, not_consider)
	} else {
		//sell
		cosider := solve_maxProfit(dp, prices, fee, i+1, 1) + prices[i] - fee
		not_consider := solve_maxProfit(dp, prices, fee, i+1, 0)
		profit = max(cosider, not_consider)
	}

	dp[i][buy] = profit

	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
