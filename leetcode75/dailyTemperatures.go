package leetcode75

func DailyTemperatures(temperatures []int) []int {

	n := len(temperatures)
	stack := make([]int, 0)
	ans := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && temperatures[i] >= temperatures[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			ans[i] = 0
		} else {
			ans[i] = stack[len(stack)-1] - i
		}

		stack = append(stack, i)
	}

	return ans
}
