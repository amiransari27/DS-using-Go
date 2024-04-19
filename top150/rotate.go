package top150

func Rotate(nums []int, k int) {

	var reverse func(l, r int) = func(l, r int) {
		for l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}

	k = k % len(nums)
	if k < 0 {
		k = k + len(nums)
	}

	reverse(0, len(nums)-k-1)
	reverse(len(nums)-k, len(nums)-1)
	reverse(0, len(nums)-1)

}
