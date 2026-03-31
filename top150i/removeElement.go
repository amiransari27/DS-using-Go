package top150i

func removeElement(nums []int, val int) int {
	p := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			//skip
		} else {
			// place number
			nums[p] = nums[i]
			p++
		}
	}

	return p
}
