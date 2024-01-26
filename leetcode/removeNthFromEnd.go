package leetcode

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	var slow, fast, pre *ListNode = head, head, nil

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next
	}
	if pre == nil {
		return head.Next
	} else {
		pre.Next = slow.Next
	}

	return head
}

func CreateList_nth(nums []int) *ListNode {
	var head, cur *ListNode

	for i := 0; i < len(nums); i++ {
		if cur == nil {
			head = &ListNode{nums[i], nil}
			cur = head
		} else {
			cur.Next = &ListNode{nums[i], nil}
			cur = cur.Next
		}
	}

	return head
}
