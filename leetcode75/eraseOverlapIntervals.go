package leetcode75

import "sort"

func EraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	i, j := 0, 1

	count := 0
	for j < len(intervals) {
		_, ce := intervals[i][0], intervals[i][1]
		ns, ne := intervals[j][0], intervals[j][1]

		if ce <= ns {
			i, j = j, j+1
		} else {
			// overlaping
			if ce <= ne {
				j = j + 1
			} else {
				i, j = j, j+1
			}
			count++
		}
	}
	return count
}
