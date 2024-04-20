package top150

func MaxProfit2(prices []int) int {
	n := len(prices)
	profit := 0

	for i := 1; i < n; i++ {
		if prices[i] > prices[i-1] {
			profit += (prices[i] - prices[i-1])
		}
	}

	return profit
}
