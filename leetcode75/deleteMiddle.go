package leetcode75

func DeleteMiddle(head *ListNode) *ListNode {

	var slow, fast *ListNode = head, head
	var pre *ListNode = nil

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		pre = slow
		slow = slow.Next
	}

	pre.Next = slow.Next

	return head
}
