package leetcode75

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func SearchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	fmt.Println(root.Val)
	if root.Val == val {
		return root
	} else if root.Val < val {
		return SearchBST(root.Left, val)
	} else {
		return SearchBST(root.Right, val)
	}
}
