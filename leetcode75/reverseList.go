package leetcode75

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil

	for head != nil {
		tmp := head
		head = head.Next
		tmp.Next = prev
		prev = tmp
		// head.Next, head, prev = prev, head.Next, head

	}
	return prev
}
