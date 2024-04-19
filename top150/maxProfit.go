package top150

func MaxProfit(prices []int) int {
	mp := 0

	l, r := 0, 1

	for r < len(prices) {
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			mp = max(mp, profit)
		} else {
			l = r
		}
		r++
	}

	return mp
}
