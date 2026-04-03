package top150i

func canJump(nums []int) bool {

	maxIndex := 0

	for i := 0; i < len(nums)-1; i++ {
		if maxIndex < i {
			return false
		}
		maxIndex = max(nums[i]+i, maxIndex)
	}

	if maxIndex >= len(nums)-1 {
		return true
	}
	return false
}
