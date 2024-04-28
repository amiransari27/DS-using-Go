package top150

func Jump(nums []int) int {
	curReach := 0
	jump := 0
	curMax := 0

	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > curMax {
			curMax = i + nums[i]
		}

		if i == curReach {
			jump++
			curReach = curMax
		}
	}

	return jump
}
