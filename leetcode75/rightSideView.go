package leetcode75

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func RightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		n := len(queue)
		var tmp *TreeNode
		for i := 0; i < n; i++ {
			tmp = queue[0]
			queue = queue[1:]
			if tmp.Left != nil {
				queue = append(queue, tmp.Left)
			}
			if tmp.Right != nil {
				queue = append(queue, tmp.Right)
			}

		}
		ans = append(ans, tmp.Val)
	}

	return ans
}
