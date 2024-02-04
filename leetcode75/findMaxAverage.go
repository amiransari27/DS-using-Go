package leetcode75

func FindMaxAverage(nums []int, k int) float64 {
	ans := 0.0
	sum := 0.0

	for i := 0; i < k; i++ {
		sum += float64(nums[i])
	}

	ans = sum / float64(k)

	for i := k; i < len(nums); i++ {
		sum += float64(nums[i]) - float64(nums[i-k])
		ans = max(ans, sum/float64(k))
	}

	return ans

}
