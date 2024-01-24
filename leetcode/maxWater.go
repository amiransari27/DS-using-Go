package leetcode

func MaxArea(height []int) int {
	low := 0
	high := len(height) - 1
	var capacity int = 0

	for low < high {
		lh := height[low]
		rh := height[high]

		if lh < rh {
			capacity = max(capacity, lh*(high-low))
			low++
		} else {
			capacity = max(capacity, rh*(high-low))
			high--
		}
	}

	return capacity
}
