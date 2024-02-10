package leetcode75

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh := MaxDepth(root.Left)
	rh := MaxDepth(root.Right)
	th := max(lh, rh) + 1
	return th
}
