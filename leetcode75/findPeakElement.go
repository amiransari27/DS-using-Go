package leetcode75

func FindPeakElement(nums []int) int {
	// for i := 0; i < len(nums)-1; i++ {
	// 	if nums[i] > nums[i+1] {
	// 		return i
	// 	}
	// }
	// return len(nums) - 1

	low, high := 0, len(nums)-1

	for low < high {
		mid := (low + high) / 2
		// mid := low + (high-low)/ 2
		if nums[mid] > nums[mid+1] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
