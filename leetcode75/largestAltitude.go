package leetcode75

func LargestAltitude(gain []int) int {
	ans := 0
	sum := 0
	for i, val := range gain {
		sum += val
		gain[i] = sum
		ans = max(ans, sum)
	}

	return ans
}
