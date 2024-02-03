package leetcode75

func MaxArea(height []int) int {
	var l, r, cap int = 0, len(height) - 1, 0

	for l < r {
		if height[l] < height[r] {
			cap = max(cap, height[l]*(r-l))
			l++
		} else {
			cap = max(cap, height[r]*(r-l))
			r--
		}
	}
	return cap
}
