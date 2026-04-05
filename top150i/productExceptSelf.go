package top150i

func productExceptSelf(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	res := make([]int, n)

	// res[i] will first store product of all elements to the left of i
	res[0] = 1
	for i := 1; i < n; i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	// Now multiply with product of all elements to the right
	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		res[i] *= rightProduct  // multiply left * right
		rightProduct *= nums[i] // update right product for next element
	}

	return res
}
