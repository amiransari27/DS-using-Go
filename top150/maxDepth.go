package top150

func MaxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}
	lh := MaxDepth(root.Left)
	rh := MaxDepth(root.Right)
	return max(lh, rh) + 1

}
