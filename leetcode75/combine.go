package leetcode

func Combine(n int, k int) [][]int {
	res := make([][]int, 0)
	tmp := []int{}
	solve_combine(&res, 1, k, tmp, n)

	return res
}

func solve_combine(res *[][]int, start int, k int, tmp []int, n int) {

	if k == 0 {
		dst := make([]int, 0)
		dst = append(dst, tmp...)
		*res = append(*res, dst)
		return
	}
	if start > n {
		return
	}

	tmp = append(tmp, start)
	solve_combine(res, start+1, k-1, tmp, n)
	tmp = tmp[:len(tmp)-1]
	solve_combine(res, start+1, k, tmp, n)
}
