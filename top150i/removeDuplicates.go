package top150i

func removeDuplicates(nums []int) int {

	val := -101
	p := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			//skip
		} else {
			//place number
			nums[p] = nums[i]
			p++
		}
		val = nums[i]
	}

	return p
}
