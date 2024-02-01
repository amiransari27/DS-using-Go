package leetcode75

func ProductExceptSelf(nums []int) []int {
	var pa, sa, res []int = make([]int, len(nums)), make([]int, len(nums)), make([]int, len(nums))
	pa[0] = 1
	sa[len(nums)-1] = 1

	for i := 1; i < len(nums); i++ {
		pa[i] = nums[i-1] * pa[i-1]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		sa[i] = nums[i+1] * sa[i+1]
	}

	for i := 0; i < len(nums); i++ {
		res[i] = sa[i] * pa[i]
	}

	return res

}
