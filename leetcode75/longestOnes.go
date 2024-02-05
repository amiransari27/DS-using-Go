package leetcode75

// [1,1,1,0,0,0,1,1,1,1,0]
func LongestOnes(nums []int, k int) int {
	ans := 0

	j := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			k--
		}
		for k < 0 {
			j++
			if nums[j] == 0 {
				k++
			}

		}
		ans = max(ans, i-j)
	}

	return ans
}
