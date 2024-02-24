package leetcode75

func FindKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, k)
}

func quickSelect(nums []int, low int, high int, k int) int {
	pi := partistionQuickSelect(nums, low, high)
	if len(nums)-k == pi {
		return nums[pi]
	} else if len(nums)-k > pi {
		return quickSelect(nums, pi+1, high, k)
	} else {
		return quickSelect(nums, low, pi-1, k)
	}
}

func partistionQuickSelect(nums []int, low int, high int) int {

	pivot := nums[high]
	position := low - 1

	for i := low; i < high; i++ {
		if nums[i] < pivot {
			position++
			nums[position], nums[i] = nums[i], nums[position]
		}
	}

	position++
	nums[position], nums[high] = nums[high], nums[position]

	return position
}
