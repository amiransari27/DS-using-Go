package top150i

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	return minDepth(root)
}
func minDepthDFS(root *TreeNode) int {

	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	var l, r int

	if root.Left != nil {
		l = minDepthDFS(root.Left)
	} else {
		l = math.MaxInt
	}

	if root.Right != nil {
		r = minDepthDFS(root.Right)
	} else {
		r = math.MaxInt
	}

	return 1 + min(l, r)

}

func minDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 1

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {

		n := len(queue)

		for n > 0 {
			front := queue[0]
			queue = queue[1:]

			if front.Left == nil && front.Right == nil {
				return res
			}

			if front.Left != nil {
				queue = append(queue, front.Left)
			}

			if front.Right != nil {
				queue = append(queue, front.Right)
			}

			n--
		}
		res++

	}

	return res
}
