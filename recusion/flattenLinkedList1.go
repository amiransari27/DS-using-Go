package recusion

type LinkList1 struct {
	Val    int
	Next   *LinkList1
	Bottom *LinkList1
}

func flattenLinkList1(head *LinkList1) *LinkList1 {

	return solveFlattenLinkList1(head)
}

func solveFlattenLinkList1(head *LinkList1) *LinkList1 {
	if head == nil {
		return nil
	}

	head2 := solveFlattenLinkList1(head.Next)

	head = mergeFlattenLinkList1_1(head, head2)

	return head
}

func mergeFlattenLinkList1_1(h1 *LinkList1, h2 *LinkList1) *LinkList1 {

	tmp := &LinkList1{
		Val:    -1,
		Next:   nil,
		Bottom: nil,
	}
	tmpTail := tmp

	for h1 != nil && h2 != nil {
		if h1.Val < h2.Val {
			tmpTail.Bottom = h1
			h1 = h1.Bottom
		} else {
			tmpTail.Bottom = h2
			h2 = h2.Bottom
		}

		tmpTail = tmpTail.Bottom
	}

	if h1 != nil {
		for h1 != nil {
			tmpTail.Bottom = h1
			h1 = h1.Bottom
			tmpTail = tmpTail.Bottom
		}
	}
	if h2 != nil {
		for h2 != nil {
			tmpTail.Bottom = h2
			h2 = h2.Bottom
			tmpTail = tmpTail.Bottom
		}
	}

	return tmp.Bottom
}

func mergeFlattenLinkList1_2(h1 *LinkList1, h2 *LinkList1) *LinkList1 {

	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}

	var tmp *LinkList1
	if h1.Val < h2.Val {
		tmp = h1
		tmp.Bottom = mergeFlattenLinkList1_2(h1.Bottom, h2)
	} else {
		tmp = h2
		tmp.Bottom = mergeFlattenLinkList1_2(h1, h2.Bottom)
	}

	return tmp
}
