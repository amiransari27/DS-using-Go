package leetcode75

type StockSpanner struct {
	stack [][]int
}

func ConstructorStockSpanner() StockSpanner {
	return StockSpanner{
		stack: make([][]int, 1),
	}
}

func (ss *StockSpanner) Next(price int) int {

	if len(ss.stack) > 0 {
		top := ss.stack[len(ss.stack)-1]
		for len(ss.stack) > 0 && top[0] < price {

		}
	} else {
		ss.stack = append(ss.stack, []int{price, 1})
	}

	return 1
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
