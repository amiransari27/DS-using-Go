package top150i

func majorityElement(nums []int) int {

	ca := nums[0]
	count := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == ca {
			count++
		} else {
			if count == 0 {
				ca = nums[i]
				count = 1
			} else {
				count--
			}
		}
	}

	return ca
}
