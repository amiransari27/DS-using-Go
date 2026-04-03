package top150i

func maxProfit(prices []int) int {

	buyPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < buyPrice {
			buyPrice = prices[i]
		} else {
			maxProfit = max(maxProfit, prices[i]-buyPrice)
		}
	}

	return maxProfit
}
