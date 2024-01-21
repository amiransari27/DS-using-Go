package leetcode

import "fmt"

func ThreeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	res := make([][]int, 0)
	memo := make(map[string]bool)
	nums = mergeSort(nums)

	for i := 0; i < len(nums)-2; i++ {
		low := i + 1
		high := len(nums) - 1

		for low < high {
			if nums[i]+nums[low]+nums[high] == 0 {
				key := fmt.Sprintf("%v%v%v", nums[i], nums[low], nums[high])

				if !memo[key] {
					memo[key] = true
					res = append(res, []int{nums[i], nums[low], nums[high]})
				}

				low++
				high--
			} else if nums[i]+nums[low]+nums[high] < 0 {
				low++
			} else {
				high--
			}
		}

	}

	return res

}

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	mid := len(nums) / 2

	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])

	ans := make([]int, 0)
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			ans = append(ans, left[0])
			left = left[1:]
		} else {
			ans = append(ans, right[0])
			right = right[1:]
		}
	}

	ans = append(ans, left...)
	ans = append(ans, right...)

	return ans
}
