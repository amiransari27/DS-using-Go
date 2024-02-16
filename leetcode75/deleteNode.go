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
func DeleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < key {
		root.Right = DeleteNode(root.Right, key)
	} else if root.Val > key {
		root.Left = DeleteNode(root.Left, key)
	} else {
		if root.Left != nil && root.Right != nil {
			val := findMax(root.Left)
			fmt.Println(val)
			root.Val = val
			root.Left = DeleteNode(root.Left, val)
		} else if root.Left != nil {
			return root.Left
		} else if root.Right != nil {
			return root.Right
		} else {
			return nil
		}
	}

	return root
}

func findMax(root *TreeNode) int {
	var maxVal int
	for root != nil {
		maxVal = root.Val
		root = root.Right
	}
	return maxVal
}
