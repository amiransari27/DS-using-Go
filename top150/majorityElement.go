package top150

func MajorityElement(nums []int) int {
	cand := nums[0]
	count := 1

	for i := 1; i < len(nums); i++ {
		if cand == nums[i] {
			count++
		} else {
			if count > 0 {
				count--
			} else {
				cand = nums[i]
				count = 1
			}
		}
	}

	return cand
}
