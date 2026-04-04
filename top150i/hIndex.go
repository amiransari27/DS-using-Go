package top150i

func hIndex(citations []int) int {

	n := len(citations)
	arr := make([]int, n+1)

	for _, c := range citations {
		if c > n {
			arr[n]++
		} else {
			arr[c]++
		}
	}

	count := 0
	for i := n; i >= 0; i-- {
		count += arr[i]
		if count >= i {
			return i
		}
	}

	return 0
}
