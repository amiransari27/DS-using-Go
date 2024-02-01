package leetcode75

func KidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := 0
	res := make([]bool, len(candies))
	for _, candie := range candies {
		if candie > maxCandies {
			maxCandies = candie
		}
	}

	for i, candie := range candies {
		if candie+extraCandies >= maxCandies {
			res[i] = true
		} else {
			res[i] = false
		}
	}
	return res
}
