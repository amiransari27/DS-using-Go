package top150i

import "math"

func containsNearbyDuplicate(nums []int, k int) bool {
	numsMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {

		if idx, exist := numsMap[nums[i]]; exist {

			if int(math.Abs(float64(idx-i))) <= k {
				return true
			}
		}
		numsMap[nums[i]] = i
	}

	return false
}
