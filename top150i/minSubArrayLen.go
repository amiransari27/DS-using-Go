package top150i

import "math"

func minSubArrayLen(target int, nums []int) int {

	minl, n := math.MaxInt, len(nums)
	sum := 0

	l, r := 0, 0

	for r < n {
		sum += nums[r]
		for sum >= target {
			minl = min(minl, r-l+1)
			sum -= nums[l]
			l++
		}

		r++
	}

	if minl == math.MaxInt {
		return 0
	}
	return minl
}
