package leetcode

import "math"

func ThreeSumClosest(nums []int, target int) int {
	quickSort_3Sum(nums)

	ans := 0
	gap := math.MaxInt

	for i := 0; i < len(nums)-2; i++ {
		low := i + 1
		high := len(nums) - 1

		for low < high {
			sum := nums[i] + nums[low] + nums[high]

			if sum < target {
				nGap := target - sum
				if nGap < gap {
					gap = nGap
					ans = sum
				}

				low++
			} else if sum > target {
				nGap := sum - target
				if nGap < gap {
					gap = nGap
					ans = sum
				}
				high--
			} else {
				return sum
			}
		}

	}

	return ans
}

func quickSort_3Sum(nums []int) []int {

	sort_3Sum(nums, 0, len(nums)-1)
	return nums
}

func sort_3Sum(nums []int, low int, high int) {
	if low < high {
		pi := partision_3Sum(nums, low, high)
		sort_3Sum(nums, low, pi-1)
		sort_3Sum(nums, pi+1, high)
	}
}

func partision_3Sum(nums []int, low int, high int) int {
	pivot := nums[high]
	position := low - 1

	for i := low; i < high; i++ {
		if nums[i] < pivot {
			position++
			tmp := nums[position]
			nums[position] = nums[i]
			nums[i] = tmp
		}
	}

	position++
	tmp := nums[position]
	nums[position] = nums[high]
	nums[high] = tmp
	return position
}
