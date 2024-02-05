package leetcode75

func LongestSubarray(nums []int) int {
	ans := 0

	j := -1
	k := 1

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
		ans = max(ans, i-j-1)
	}

	return ans
}
