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
	return k + 1
}

func RemoveDuplicates2(nums []int) int {

	k := 0
	curVal := nums[0]
	curCount := 1
	for i := 1; i < len(nums); i++ {
		if curVal != nums[i] {
			k++
			nums[k] = nums[i]
			curVal = nums[k]
			curCount = 1
		} else if curCount < 2 {
			k++
			nums[k] = nums[i]
			curCount++
		}
	}
	return k + 1
}
