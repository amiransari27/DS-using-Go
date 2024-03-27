package leetcode75

func CountBits(n int) []int {
	res := make([]int, n+1)

	res[0] = 0

	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			res[i] = res[i/2]
		} else {
			res[i] = res[i/2] + 1
		}
	}

	return res
}
