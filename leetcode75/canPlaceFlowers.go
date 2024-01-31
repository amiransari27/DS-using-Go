package leetcode75

func CanPlaceFlowers(flowerbed []int, n int) bool {

	idx := 0

	for n > 0 && idx < len(flowerbed) {
		var (
			left  int
			right int
		)
		if idx == 0 {
			left = 0
		} else {
			left = flowerbed[idx-1]
		}
		if idx == len(flowerbed)-1 {
			right = 0
		} else {
			right = flowerbed[idx+1]
		}

		if flowerbed[idx] == 0 && left == 0 && right == 0 {
			flowerbed[idx] = 1
			n--
		}
		idx++
	}

	return n == 0
}
