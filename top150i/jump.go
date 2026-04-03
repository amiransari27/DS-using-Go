package top150i

import "math"

func jump(nums []int) int {
	n := len(nums)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}
	return solveJump(nums, n, 0, memo)
}

func solveJump(nums []int, n int, i int, memo []int) int {
	if i >= n-1 {
		return 0
	}
	if memo[i] != -1 {
		return memo[i]
	}
	minJumps := math.MaxInt32
	for j := 1; j <= nums[i]; j++ {
		jumps := 1 + solveJump(nums, n, i+j, memo)
		if jumps < minJumps {
			minJumps = jumps
		}
	}
	memo[i] = minJumps
	return minJumps
}

func jumpBU(nums []int) int {

	n := len(nums)

	dp := make([]int, n)

	for i := range dp {
		dp[i] = math.MaxInt
	}

	dp[0] = 0

	for i := 1; i < n; i++ {

		for j := 0; j < i; j++ {

			if j+nums[j] >= i {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}

	}

	return dp[n-1]
}
