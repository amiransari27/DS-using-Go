package top150

func canJump(nums []int) bool {

	// memo := make([]int, 10001)
	// for i := range memo {
	// 	memo[i] = -1
	// }

	// return solveCanJump(memo, nums, 0)

	dp := make([]bool, len(nums))
	dp[0] = true

	for i := 1; i < len(nums); i++ {
		j := i - 1
		for j >= 0 {
			if dp[j] && j+nums[j] >= i {
				dp[i] = true
				break
			}
            j--
		}

	}

	return dp[len(nums)-1]
}

func solveCanJump(memo []int, nums []int, i int) bool {

	if i == len(nums)-1 {
		return true
	}

	if memo[i] != -1 {
		if memo[i] == 0 {
			return false
		} else {
			return true
		}
	}

	ml := nums[i]

	for k := 1; k <= ml; k++ {
		if solveCanJump(memo, nums, i+k) {
			memo[i] = 1
			return true
		}

	}

	memo[i] = 0
	return false
}
