package segmenttree

func build(nums []int) []int {
	n := len(nums)
	res := make([]int, 4*n)

	solveBuild(0, 0, n-1, &res, nums)
	return res
}

func solveBuild(i int, l int, r int, arr *[]int, num []int) {

	if l == r {
		(*arr)[i] = num[l]
		return
	}

	mid := (l + r) / 2
	// left call
	solveBuild(2*i+1, l, mid, arr, num)
	// right call
	solveBuild(2*i+2, mid+1, r, arr, num)

	(*arr)[i] = (*arr)[2*i+1] + (*arr)[2*i+2]
}
