package leetcode75

func CombinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	solve_combinationSum3(&res, 1, k, n, []int{})
	return res
}

func solve_combinationSum3(res *[][]int, start int, k int, n int, tmp []int) {

	if k == 0 {
		dst := make([]int, 0)
		dst = append(dst, tmp...)
		sum := 0
		for _, val := range dst {
			sum += val
		}
		if sum == n {
			*res = append(*res, dst)
		}
		return
	}
	if start > 9 {
		return
	}

	tmp = append(tmp, start)
	solve_combinationSum3(res, start+1, k-1, n, tmp)
	tmp = tmp[:len(tmp)-1]
	solve_combinationSum3(res, start+1, k, n, tmp)

}
