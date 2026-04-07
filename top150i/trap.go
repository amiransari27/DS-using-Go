package top150i

func trap(height []int) int {

	n := len(height)
	lm := make([]int, n)
	rm := make([]int, n)

	lm[0] = height[0]
	for i := 1; i < n; i++ {
		lm[i] = max(lm[i-1], height[i])
	}

	rm[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rm[i] = max(rm[i+1], height[i])
	}

	trapWater := 0
	for i := 0; i < n; i++ {
		trapWater += min(lm[i], rm[i]) - height[i]
	}

	return trapWater
}
