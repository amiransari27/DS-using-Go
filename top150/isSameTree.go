package top150

func IsSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)

}

func IsSameTreeBFS(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	q1 := make([]*TreeNode, 0)
	q2 := make([]*TreeNode, 0)

	q1 = append(q1, p)
	q2 = append(q2, q)

	for len(q1) > 0 && len(q2) > 0 {
		n1 := q1[len(q1)-1]
		q1 = q1[:len(q1)-1]

		n2 := q2[len(q2)-1]
		q2 = q2[:len(q2)-1]

		if n1.Val != n2.Val {
			return false
		}

		if n1.Left != nil && n2.Left != nil {
			q1 = append(q1, n1.Left)
			q2 = append(q2, n2.Left)
		} else if n1.Left != nil || n2.Left != nil {
			return false
		}

		if n1.Right != nil && n2.Right != nil {
			q1 = append(q1, n1.Right)
			q2 = append(q2, n2.Right)
		} else if n1.Right != nil || n2.Right != nil {
			return false
		}
	}

	return true
}
