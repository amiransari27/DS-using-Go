package leetcode75

import (
	"strconv"
	"strings"
)

func LeafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	r1Arr := []string{}
	r2Arr := []string{}
	leafs(root1, &r1Arr)
	leafs(root2, &r2Arr)

	return strings.Join(r1Arr, "-") == strings.Join(r2Arr, "-")
}

func leafs(root *TreeNode, leafArr *[]string) {

	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*leafArr = append(*leafArr, strconv.Itoa(root.Val))
	}

	leafs(root.Left, leafArr)
	leafs(root.Right, leafArr)
}
