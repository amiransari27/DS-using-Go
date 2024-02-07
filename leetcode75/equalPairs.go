package leetcode75

import (
	"fmt"
	"strconv"
	"strings"
)

func EqualPairs(grid [][]int) int {

	l := len(grid)
	count := 0

	for r := 0; r < l; r++ {
		for c := 0; c < l; c++ {
			isEqual := true
			for i := 0; i < l; i++ {
				if grid[r][i] != grid[i][c] {
					isEqual = false
					break
				}
			}

			if isEqual {
				count++
			}
		}
	}

	return count
}

func EqualPairsOptimize(grid [][]int) int {
	l := len(grid)
	count := 0

	rmap := map[string]int{}
	for r := 0; r < l; r++ {
		keyArr := []string{}
		for c := 0; c < l; c++ {
			keyArr = append(keyArr, strconv.Itoa(grid[r][c]))
		}
		key := strings.Join(keyArr, "-")
		if rmap[key] == 0 {
			rmap[key] = 1
		} else {
			rmap[key] += 1
		}
	}

	fmt.Println(rmap)

	for c := 0; c < l; c++ {
		keyArr := []string{}
		for r := 0; r < l; r++ {
			keyArr = append(keyArr, strconv.Itoa(grid[r][c]))
		}
		key := strings.Join(keyArr, "-")
		if rmap[key] > 0 {
			count += rmap[key]
		}
	}

	return count
}
