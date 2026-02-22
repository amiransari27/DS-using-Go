package dp

import "sort"

func largestDivisibleSubset(nums []int) []int {

	sort.Ints(nums)

	var result []int

	var solve func(int, int, []int)

	solve = func(idx, pre int, temp []int) {

		if idx >= len(nums) {
			if len(temp) > len(result) {
				result = append([]int{}, temp...)
			}

			return
		}
		//take
		if pre == -1 || nums[idx]%pre == 0 {
			temp = append(temp, nums[idx])
			solve(idx+1, nums[idx], temp)
			temp = temp[:len(temp)-1]
		}

		solve(idx+1, pre, temp)
	}

	solve(0, -1, make([]int, 0))

	return result
}
