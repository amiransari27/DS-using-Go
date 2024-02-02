package leetcode75

func MoveZeroes(nums []int) {
	i, j := 0, 0

	for j < len(nums) {
		if nums[i] != 0 {
			i++
			j++
			continue
		}
		if nums[j] == 0 {
			j++
			continue
		}
		tmp := nums[j]
		nums[j] = nums[i]
		nums[i] = tmp
		i++
		j++
	}
}
