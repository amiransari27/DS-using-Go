package top150i

func candy(ratings []int) int {

	n := len(ratings)
	l2r := make([]int, n)
	r2l := make([]int, n)

	for i := 0; i < n; i++ {
		l2r[i], r2l[i] = 1, 1
	}

	// check left to right

	for i := 0; i < n; i++ {
		if i > 0 && ratings[i-1] < ratings[i] {
			l2r[i] = l2r[i-1] + 1
		}
	}

	// now check r2l
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && ratings[i] > ratings[i+1] {
			r2l[i] = r2l[i+1] + 1
		}
	}

	total := 0

	for i := 0; i < n; i++ {
		if l2r[i] > r2l[i] {
			total += l2r[i]
		} else {
			total += r2l[i]
		}
	}

	return total
}
