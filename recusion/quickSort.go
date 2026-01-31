package recusion

func QuickSort(nums []int, low int, high int) {

	if low >= high {
		return
	}

	p := partision(nums, low, high)

	QuickSort(nums, low, p-1)
	QuickSort(nums, p+1, high)
}

func partision(nums []int, low int, high int) int {

	i, j := low, low

	for i < high {
		if nums[i] < nums[high] {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
		i++
	}
	nums[j], nums[high] = nums[high], nums[j]
	return j
}
