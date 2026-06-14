package top150i

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {

}

func maxDepthDFS(root *TreeNode) int {

	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}
	var l, r int

	if root.Left != nil {
		l = maxDepthDFS(root.Left)
	}
	if root.Right != nil {
		r = maxDepthDFS(root.Right)
	}

	return 1 + max(l, r)
}

func maxDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	depth := 1
	for len(queue) > 0 {
		n := len(queue)

		for n > 0 {
			front := queue[0]
			queue = queue[1:]

			if front.Left != nil {
				queue = append(queue, front.Left)
			}
			if front.Right != nil {
				queue = append(queue, front.Right)
			}

			n--
		}
		depth++
	}

	return depth
}
