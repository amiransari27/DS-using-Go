package problems

import "fmt"

type GTNode struct {
	val      int
	children []*GTNode
}

func ConstructGenericTree() *GTNode {
	intput := []int{10, 20, 50, -1, 60, -1, -1, 30, 70, -1, 80, 110, -1, 120, -1, -1, 90, -1, -1, 40, 100, -1, -1, -1}

	root := GTNode{
		val: intput[0],
	}

	stack := []*GTNode{}
	stack = append(stack, &root)

	idx := 1
	for len(stack) > 0 && idx < len(intput) {
		if intput[idx] == -1 {
			// pop from the stack
			fmt.Println("--")
			stack = stack[:len(stack)-1]
		} else {
			//create node
			node := GTNode{
				val: intput[idx],
			}
			fmt.Print(node)
			// asign it to last node in the stack
			lastNode := stack[len(stack)-1]
			fmt.Println(lastNode)
			lastNode.children = append(lastNode.children, &node)
			// push node to the stack
			fmt.Println(len(lastNode.children))
			stack = append(stack, &node)
		}

		idx++
	}

	return &root
}
