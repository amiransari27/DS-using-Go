package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// add two number .
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var head, cur *ListNode = nil, nil

	var carry int = 0

	for l1 != nil && l2 != nil {
		var v1, v2 int = 0, 0
		if l1 != nil {
			v1 = l1.Val
		}
		if l2 != nil {
			v2 = l2.Val
		}

		sum := v1 + v2 + carry
		carry = sum / 10

		newNode := &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		if cur != nil {
			cur.Next = newNode
			cur = cur.Next
		} else {
			head = newNode
			cur = newNode
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	if carry > 0 {
		newNode := &ListNode{
			Val:  carry,
			Next: nil,
		}
		cur.Next = newNode
		cur = cur.Next
	}

	return head
}

func CreateLinkList(nums []int) *ListNode {

	var head, pre, cur *ListNode = nil, nil, nil

	for i := 0; i < len(nums); i++ {
		newNode := &ListNode{
			Val:  nums[i],
			Next: nil,
		}
		if pre == nil {
			head = newNode
			pre = newNode
			cur = newNode
			continue
		}
		cur = newNode
		pre.Next = cur
		pre = cur
	}

	return head
}

func PrintLinkList(l1 *ListNode) {
	for l1 != nil {
		fmt.Println(l1)
		l1 = l1.Next
	}
}
