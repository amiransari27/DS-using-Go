package leetcode75

func PairSum(head *ListNode) int {
	var fast, slow *ListNode = head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var revHead *ListNode = nil
	curr := slow

	for curr != nil {
		tmp := curr
		curr = curr.Next
		tmp.Next = revHead
		revHead = tmp
	}

	maxVal := 0

	for revHead != nil {
		maxVal = max(maxVal, head.Val+revHead.Val)
		head = head.Next
		revHead = revHead.Next
	}

	return maxVal
}
