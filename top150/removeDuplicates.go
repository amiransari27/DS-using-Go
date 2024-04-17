package top150

import "math"

func RemoveDuplicates(nums []int) int {
	k := -1

	curVal := math.MinInt
	for i := 0; i < len(nums); i++ {
		if curVal != nums[i] {
			k++
			nums[k] = nums[i]
			curVal = nums[k]
		}
	}
	return k
}
