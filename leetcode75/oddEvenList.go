package leetcode75

func OddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var even, ec, odd, oc *ListNode

	idx := -1
	for head != nil {
		idx++

		if idx%2 == 0 {
			if ec == nil {
				even = head
				ec = even
			} else {
				ec.Next = head
				ec = ec.Next
			}
		} else {
			if oc == nil {
				odd = head
				oc = odd
			} else {
				oc.Next = head
				oc = oc.Next
			}
		}

		head = head.Next
	}
	ec.Next = odd
	oc.Next = nil
	return even
}
