package recusion

func ReverseStrack(stack []int) []int {

	reverse(&stack)

	return stack
}

func reverse(stack *[]int) {

	if len(*stack) == 0 {
		return
	}

	top := (*stack)[len(*stack)-1]
	(*stack) = (*stack)[:len(*stack)-1]

	reverse(stack)

	insertAtBottom(stack, top)

	// (*stack) = append([]int{top}, (*stack)...)
}

func insertAtBottom(stack *[]int, ele int) {

	if len(*stack) == 0 {
		(*stack) = append((*stack), ele)
		return
	}

	top := (*stack)[len(*stack)-1]
	(*stack) = (*stack)[:len(*stack)-1]

	insertAtBottom(stack, ele)
	(*stack) = append((*stack), top)
}
