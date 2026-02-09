package recusion

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flattenBST(root *TreeNode) *TreeNode {

	if root == nil {
		return root
	}

	head := flattenBST(root.Left)
	root.Left = nil

	if head == nil {
		head = root
	} else {
		temp := head
		for temp != nil && temp.Right != nil {
			temp = temp.Right
		}
		temp.Right = root
	}

	root.Right = flattenBST(root.Right)

	return head

}
