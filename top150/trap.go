package top150

func Trap(height []int) int {
	n := len(height)

	lMax := make([]int, n)
	lMax[0] = height[0]
	for i := 1; i < n; i++ {
		lMax[i] = max(lMax[i-1], height[i])
	}
	rMax := make([]int, n)
	rMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rMax[i] = max(rMax[i+1], height[i])
	}

	sum := 0

	for i := 0; i < n; i++ {
		h := min(lMax[i], rMax[i]) - height[i]
		sum += h
	}

	return sum
}
