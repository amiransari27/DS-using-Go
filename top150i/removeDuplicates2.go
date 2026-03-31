package top150i

func removeDuplicates2(nums []int) int {

	p := 0
	count := 0
	val := -10001

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			count = 0
		}
		if nums[i] != val || (nums[i] == val && count < 2) {
			//place
			nums[p] = nums[i]
			p++
			count++
		}

		val = nums[i]
	}

	return p
}
