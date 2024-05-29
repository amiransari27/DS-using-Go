package top150

import (
	"fmt"
	"slices"
)

func ThreeSum(nums []int) [][]int {
	slices.Sort(nums)
	ans := make([][]int, 0)
	memo := make(map[string]bool)
	n := len(nums)
	for i := 0; i <= n-3; i++ {
		l, r := i+1, n-1

		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				key := fmt.Sprintf("%v%v%v", nums[i], nums[l], nums[r])

				if memo[key] {

				} else {
					memo[key] = true
					ans = append(ans, []int{nums[i], nums[l], nums[r]})
				}
			} else if sum < 0 {
				l++
				continue
			} else {
				r--
				continue
			}

			l++
			r--
		}
	}

	return ans
}
