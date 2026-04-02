func reverse(nums []int, start int, end int) {

	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
func rotate(nums []int, k int) {

	k = k % len(nums)

	if k == 0 {
		return
	}

	//reverse all
	reverse(nums, 0, len(nums)-1)

	//reverse till k-1

	reverse(nums, 0, k-1)

	//reverse from  k to len -q
	reverse(nums, k, len(nums)-1)
} 