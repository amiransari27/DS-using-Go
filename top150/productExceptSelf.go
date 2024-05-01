package top150

import "fmt"

func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	var pa, sa, res []int = make([]int, n), make([]int, n), make([]int, n)
	pa[0] = 1
	sa[len(nums)-1] = 1

	for i := 1; i < n; i++ {
		pa[i] = nums[i-1] * pa[i-1]
	}

	fmt.Println(pa)

	for i := n - 2; i >= 0; i-- {
		sa[i] = nums[i+1] * sa[i+1]
	}

	fmt.Println(sa)

	for i := 0; i < n; i++ {
		res[i] = sa[i] * pa[i]
	}

	return res
}
