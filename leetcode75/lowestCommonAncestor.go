package leetcode75

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	leftN := LowestCommonAncestor(root.Left, p, q)
	rightN := LowestCommonAncestor(root.Right, p, q)

	if leftN != nil && rightN != nil {
		return root
	}

	if leftN == nil {
		return rightN
	}
	return rightN
}
