package leetcode75

import "fmt"

func OrangesRotting(grid [][]int) int {
	direction := [4][2]int{
		{0, -1}, {-1, 0}, {0, 1}, {1, 0},
	}
	total, rotten, time := 0, 0, 0
	queue := make([][]int, 0)

	for i, vi := range grid {
		for j, jv := range vi {

			if jv == 2 || jv == 1 {
				total++
			}
			if jv == 2 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	fmt.Println(total)

	if total == 0 {
		return 0
	}
	for len(queue) > 0 {
		size := len(queue)
		rotten += size
		if rotten == size {
			return time
		}
		for size > 0 {
			pair := queue[0]
			queue = queue[1:]
			i := pair[0]
			j := pair[1]

			for _, dir := range direction {
				new_i := i + dir[0]
				new_j := j + dir[1]

				if new_i >= 0 && new_i < len(grid) && new_j >= 0 && new_j < len(grid[0]) && grid[new_i][new_j] == 1 {
					queue = append(queue, []int{new_i, new_j})
					grid[new_i][new_j] = 2
				}
			}

			size--
		}

		time++
	}

	return -1
}
