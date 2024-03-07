package leetcode75

import "slices"

func MinEatingSpeed(piles []int, h int) int {
	low := 1
	high := slices.Max(piles)

	for low < high {
		mid := (low + high) / 2

		if canEatAllBananas(piles, mid, h) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func canEatAllBananas(piles []int, mid int, h int) bool {
	totalHours := 0

	for _, b := range piles {
		totalHours += b / mid
		if b%mid > 0 {
			totalHours++
		}
	}

	return totalHours <= h
}
