package leetcode75

import "fmt"

func Rob(nums []int) int {

	memo := make([]int, 101)
	for i, _ := range memo {
		memo[i] = -1
	}

	return solve_rob(nums, 0, memo)
}

func solve_rob(nums []int, i int, memo []int) int {

	if i >= len(nums) {
		fmt.Println(i)
		return 0
	}

	if memo[i] != -1 {
		return memo[i]
	}

	steal := nums[i] + solve_rob(nums, i+2, memo)
	skip := solve_rob(nums, i+1, memo)

	memo[i] = max(steal, skip)

	return memo[i]
}
