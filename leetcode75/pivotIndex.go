package leetcode75

func PivotIndex(nums []int) int {
	sumArr := make([]int, len(nums))
	sumArr[0] = 0
	sum := nums[0]

	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		sumArr[i] = nums[i-1] + sumArr[i-1]
	}

	for i := 0; i < len(nums); i++ {
		if sumArr[i] == sum-nums[i]-sumArr[i] {
			return i
		}
	}

	return -1
}
