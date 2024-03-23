package leetcode

import "math"

func MaxSubArray(nums []int) int {

	n := len(nums)
	maxSum := math.MinInt

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += nums[k]
			}

			maxSum = max(maxSum, sum)
		}
	}

	return maxSum
}
func MaxSubArray2(nums []int) int {

	n := len(nums)
	maxSum := math.MinInt

	for i := 0; i < n; i++ {
		currSum := 0
		for j := i; j < n; j++ {
			currSum += nums[j]
			maxSum = max(maxSum, currSum)
		}
	}

	return maxSum
}

func MaxSubArray3(nums []int) int {

	n := len(nums)
	maxSum := math.MinInt
	currSum := 0

	for i := 0; i < n; i++ {
		currSum += nums[i]
		maxSum = max(maxSum, currSum)
		if currSum < 0 {
			currSum = 0
		}
	}

	return maxSum
}
