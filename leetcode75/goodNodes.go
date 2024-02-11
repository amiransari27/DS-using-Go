package leetcode75

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func GoodNodes(root *TreeNode) int {
	return goodNodeSol(root, root.Val)
}

func goodNodeSol(root *TreeNode, maxVal int) int {
	if root == nil {
		return 0
	}

	maxVal = max(maxVal, root.Val)
	c1 := goodNodeSol(root.Left, maxVal)
	c2 := goodNodeSol(root.Right, maxVal)

	if root.Val >= maxVal {
		return c1 + c2 + 1
	}

	return c1 + c2
}
