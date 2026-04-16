package top150i

import (
	"slices"
)

func threeSum(nums []int) [][]int {

	slices.Sort(nums)
	n := len(nums)
	res := make([][]int, 0)

	if n < 3 {
		return res
	}

	for i := 0; i <= n-3; i++ {
		// Skip duplicate values for i
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := -nums[i]
		l, r := i+1, n-1

		for l < r {

			if nums[l]+nums[r] < target {
				l++
			} else if nums[l]+nums[r] > target {
				r--
			} else {

				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}

				res = append(res, []int{nums[i], nums[l], nums[r]})

				l++
				r--

			}

		}

	}

	return res
}
