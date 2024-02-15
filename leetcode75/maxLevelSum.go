package leetcode75

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func MaxLevelSum(root *TreeNode) int {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	level := 0
	var ans int
	maxSum := math.MinInt

	for len(queue) > 0 {
		level++
		n := len(queue)
		tmpSum := 0
		for n > 0 {
			node := queue[0]
			queue = queue[1:]

			tmpSum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			n--
		}

		if tmpSum > maxSum {
			maxSum = tmpSum
			ans = level
		}
	}

	return ans
}
