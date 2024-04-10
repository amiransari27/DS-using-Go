package top150

func RemoveElement(nums []int, val int) int {
	k := 0

	for i, v := range nums {
		if v != val {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
