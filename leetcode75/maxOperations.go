package leetcode75

import "slices"

func MaxOperations(nums []int, k int) int {
	slices.Sort(nums)

	var low, high = 0, len(nums) - 1
	count := 0
	for low < high {
		if nums[low]+nums[high] == k {
			count++
			low++
			high--
		} else if nums[low]+nums[high] < k {
			low++
		} else {
			high--
		}
	}
	return count
}
