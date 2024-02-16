package problems

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// []int{5, 3, 2, 0, 0, 4, 0, 0, 7, 6, 0, 0, 10, 8, 0, 0, 0}
func Deserialize(arr *[]int) *TreeNode {
	if len(*arr) == 0 {
		return nil
	}
	tmp := (*arr)[0]
	*arr = (*arr)[1:]
	if tmp == 0 {
		return nil
	}

	root := &TreeNode{
		Val: tmp,
	}
	root.Left = Deserialize(arr)
	root.Right = Deserialize(arr)

	return root
}

func Serialize(root *TreeNode) []int {
	if root == nil {
		return []int{0}
	}

	ans := []int{root.Val}

	left := Serialize(root.Left)
	rigth := Serialize(root.Right)

	ans = append(ans, left...)
	ans = append(ans, rigth...)
	return ans
}

// func main() {

// 	arr := []int{5, 3, 2, 0, 0, 4, 0, 0, 7, 6, 0, 0, 10, 8, 0, 0, 0}

// 	root := problems.Deserialize(&arr)
// 	fmt.Printf("%+v\n", root.Left.Right)

// 	ans := problems.Serialize(root)
// 	fmt.Printf("%+v\n", ans)

// }
