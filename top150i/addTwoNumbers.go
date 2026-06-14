package top150i

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	carry := 0
	ans := &ListNode{}
	curr := ans

	for l1 != nil || l2 != nil {
		var a, b int = 0, 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}

		sum := a + b + carry
		carry = sum / 10

		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next

	}
	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}

	return ans.Next
}
