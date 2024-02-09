package leetcode75

type Node struct {
	Val  byte
	Next *Node
}

func PredictPartyVictory(senate string) string {
	var Rcount, Dcount int
	var head, curr *Node

	for i := 0; i < len(senate); i++ {
		if senate[i] == 'R' {
			Rcount++
		} else {
			Dcount++
		}
		if head == nil {
			head = &Node{
				Val:  senate[i],
				Next: nil,
			}
			curr = head
		} else {
			curr.Next = &Node{
				Val:  senate[i],
				Next: nil,
			}
			curr = curr.Next
		}
	}
	// creating circular list
	curr.Next = head

	for Rcount > 0 && Dcount > 0 {
		if head.Val == 'R' {
			removeNext(head, 'D')
			Dcount--
		} else {
			removeNext(head, 'R')
			Rcount--
		}
		head = head.Next
	}

	if Rcount == 0 {
		return "Dire"
	} else {
		return "Radiant"
	}
}

func removeNext(head *Node, char byte) {

	curr := head
	for {
		if curr.Next.Val == char {
			curr.Next = curr.Next.Next
			break
		}

		curr = curr.Next
	}
}
