package leetcode75

import (
	"math"
	"slices"
)

func SuccessfulPairs(spells []int, potions []int, success int64) []int {

	// spells * minPortion >= success
	// minPortion >= ceil(success/spell)
	n := len(potions)
	res := []int{}
	slices.Sort[[]int](potions)

	for _, spell := range spells {
		minPor := math.Ceil(float64(success) / float64(spell))
		pos, _ := slices.BinarySearch(potions, int(minPor))
		res = append(res, n-pos)
	}

	return res
}
