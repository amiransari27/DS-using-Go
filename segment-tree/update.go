package segmenttree

func update(idx int, val int, i int, l int, r int, st *[]int, num []int) {

	if l == r {
		num[idx] = val
		(*st)[i] = val
		return
	}

	mid := (l + r) / 2

	if idx <= mid {
		update(idx, val, 2*i+1, l, mid, st, num)
	} else {
		update(idx, val, 2*i+2, mid+1, r, st, num)
	}

	(*st)[i] = (*st)[2*i+1] + (*st)[2*i+2]

}
