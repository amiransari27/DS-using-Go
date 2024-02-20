package leetcode75

func NearestExit(maze [][]byte, entrance []int) int {
	direction := [4][2]int{
		{0, -1}, {-1, 0}, {0, 1}, {1, 0},
	}
	m := len(maze)
	n := len(maze[0])

	queue := make([][]int, 0)
	queue = append(queue, []int{entrance[0], entrance[1]})
	// make visited by making it a wall '+'
	maze[entrance[0]][entrance[1]] = '+'

	step := 0

	for len(queue) > 0 {
		N := len(queue)

		for N > 0 {
			pair := queue[0]
			queue = queue[1:]
			i := pair[0]
			j := pair[1]

			if (i != entrance[0] || j != entrance[1]) &&
				(i == 0 || i == m-1 || j == 0 || j == n-1) {
				return step
			}

			// explore all neighbours
			for _, dir := range direction {
				new_i := i + dir[0]
				new_j := j + dir[1]

				if new_i >= 0 && new_i < m && new_j >= 0 && new_j < n && maze[new_i][new_j] != '+' {
					queue = append(queue, []int{new_i, new_j})
					maze[new_i][new_j] = '+'
				}
			}

			N--
		}
		step++
	}

	return -1
}
