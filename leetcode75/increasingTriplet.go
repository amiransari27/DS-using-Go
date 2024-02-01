package leetcode75

import "math"

func IncreasingTriplet(nums []int) bool {
	i := math.MaxInt32
	j := math.MaxInt32

	for _, val := range nums {
		if val <= i {
			i = val
		} else if val <= j {
			j = val
		} else {
			return true
		}
	}
	return false
}
