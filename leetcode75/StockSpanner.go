package leetcode75

type StockSpanner struct {
	stack [][]int
}

func ConstructorStockSpanner() StockSpanner {
	return StockSpanner{
		stack: make([][]int, 0),
	}
}

func (ss *StockSpanner) Next(price int) int {

	span := 1

	for len(ss.stack) > 0 && ss.stack[len(ss.stack)-1][0] <= price {

		span += ss.stack[len(ss.stack)-1][1]

		ss.stack = ss.stack[:len(ss.stack)-1]
	}

	ss.stack = append(ss.stack, []int{price, span})

	return span
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
