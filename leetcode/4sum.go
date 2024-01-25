package leetcode

func FourSum(nums []int, target int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	// sort.Ints(nums)
	nums = mergeSort_fourSum(nums)
	ans := make([][]int, 0)

	//-2,-1,-1,1,1,2,2
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			low := j + 1
			high := len(nums) - 1

			for low < high {
				sum := nums[i] + nums[j] + nums[low] + nums[high]
				if sum == target {
					sumArr := []int{nums[i], nums[j], nums[low], nums[high]}
					ans = append(ans, sumArr)

					for low < high && nums[low] == sumArr[2] {
						low++
					}
					for low < high && nums[high] == sumArr[3] {
						high--
					}

				} else if sum < target {
					low++
				} else {
					high--
				}
			}
		}
	}

	return ans
}

func mergeSort_fourSum(nums []int) []int {

	if len(nums) < 2 {
		return nums
	}

	mid := len(nums) / 2
	left := mergeSort_fourSum(nums[:mid])
	right := mergeSort_fourSum(nums[mid:])

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
