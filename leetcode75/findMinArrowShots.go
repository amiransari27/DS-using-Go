package leetcode75

import "sort"

func FindMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	stack := make([][]int, 0)
	stack = append(stack, points[0])

	for i := 1; i < len(points); i++ {
		//pick top element from stack
		top := stack[len(stack)-1]
		// pop element
		stack = stack[:len(stack)-1]

		cs, ce := top[0], top[1]
		ns, ne := points[i][0], points[i][1]

		if ce < ns {
			// no overlapping
			stack = append(stack, top)
			stack = append(stack, points[i])
		} else {
			if ce <= ne {
				stack = append(stack, []int{cs, ce})
			} else {
				stack = append(stack, []int{cs, ne})
			}

		}

	}

	return len(stack)
}
