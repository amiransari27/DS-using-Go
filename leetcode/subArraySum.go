package leetcode

func SubArraySum(arr []int, target int) int {

	currSum := 0
	result := 0
	memo := map[int]bool{}
	memo[0] = true
	for _, val := range arr {
		currSum += val
		diff := currSum - target
		if memo[diff] {
			result++
		} else {
			memo[currSum] = true
		}

	}

	return result
}
