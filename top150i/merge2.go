package top150i

import (
	"slices"
)

func merge2(intervals [][]int) [][]int {

	res := make([][]int, 0)

	slices.SortFunc(intervals, func(a []int, b []int) int {
		return a[0] - b[0]
	})

	for i := 0; i < len(intervals); {
		j := i + 1
		maxSecond := intervals[i][1]

		for j < len(intervals) && maxSecond >= intervals[j][0] {
			j++

			if maxSecond < intervals[j-1][1] {
				maxSecond = intervals[j-1][1]
			}
		}

		res = append(res, []int{intervals[i][0], maxSecond})

		i = j
	}

	return res
}
