package top150i

func maxArea(height []int) int {

	l, r := 0, len(height)-1

	ans := 0

	for l < r {
		h := height[l]
		w := r - l

		if height[r] < h {
			h = height[r]
			r--
		} else {
			l++
		}

		ans = max(ans, h*w)
	}

	return ans
}
