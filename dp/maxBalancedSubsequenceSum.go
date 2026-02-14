// 2926. Maximum Balanced Subsequence Sum
package dp

import "math"

func maxBalancedSubsequenceSum(nums []int) int64 {
	maxVal := int64(math.MinInt)

	for _, v := range nums {
		maxVal = max(maxVal, int64(v))
	}
	if maxVal <= 0 {
		return maxVal
	}

	memo := make(map[[2]int]int64)

	return solveMaxBalancedSubsequenceSum(0, -1, nums, memo)
}

func solveMaxBalancedSubsequenceSum(i int, pi int, nums []int, memo map[[2]int]int64) int64 {

	if i == len(nums) {
		return 0
	}

	key := [2]int{i, pi}
	if pi != -1 {
		if val, exist := memo[key]; exist {
			return val
		}
	}

	take := int64(math.MinInt)

	not_take := int64(math.MinInt)
	if pi == -1 || nums[i]-i >= nums[pi]-pi {
		take = int64(nums[i]) + solveMaxBalancedSubsequenceSum(i+1, i, nums, memo)
	}

	not_take = solveMaxBalancedSubsequenceSum(i+1, pi, nums, memo)

	if pi != -1 {
		memo[key] = max(take, not_take)
	}

	return max(take, not_take)
}

func maxBalancedSubsequenceSumBU(nums []int) int64 {
	maxVal := int64(math.MinInt)

	for _, v := range nums {
		maxVal = max(maxVal, int64(v))
	}
	if maxVal <= 0 {
		return maxVal
	}

	var result int64 = math.MinInt64
	n := len(nums)

	t := make([]int64, n)

	for i := 0; i < n; i++ {
		// Base case: take just nums[i] by itself
		t[i] = int64(nums[i])
		result = max(result, t[i])

		// Try extending from previous elements that satisfy the condition
		for j := 0; j < i; j++ {
			if nums[i]-i >= nums[j]-j {
				t[i] = max(t[i], t[j]+int64(nums[i]))
				result = max(result, t[i])
			}
		}
	}

	return result
}
