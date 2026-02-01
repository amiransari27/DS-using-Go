package dp

type IntList []int

func Rob2(nums []int) int {

	n := len(nums)
	t := make(IntList, len(nums))
	t.fillWith(-1)

	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	withTakingZeroIndex := solveRob2(nums, 0, n-2, t)
	t.fillWith(-1)
	withoutTakingZeroIndex := solveRob2(nums, 1, n-1, t)

	return max(withTakingZeroIndex, withoutTakingZeroIndex)

}

func solveRob2(nums []int, l int, r int, t IntList) int {

	if l > r {
		return 0
	}

	if t[l] != -1 {
		return t[l]
	}

	steal := nums[l] + solveRob2(nums, l+2, r, t)
	notSteal := solveRob2(nums, l+1, r, t)

	t[l] = max(steal, notSteal)

	return t[l]

}

func Rob2BU(nums []int) int {

	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}

	withTakingZeroIndex := solveRob2BU(nums, 0, n-2)

	withoutTakingZeroIndex := solveRob2BU(nums, 1, n-1)

	return max(withTakingZeroIndex, withoutTakingZeroIndex)
}

func solveRob2BU(nums []int, l int, r int) int {

	t := make([]int, r-l+1)

	t[0] = nums[l]
	t[1] = max(nums[l], nums[l+1])
	for k := 2; k < len(t); k++ {
		steal := nums[l+k] + t[k-2]
		notSteal := t[k-1]
		t[k] = max(steal, notSteal)
	}

	return t[len(t)-1]
}

func Rob2BUConstSpace(nums []int) int {

	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}

	withTakingZeroIndex := solveRob2ConstSpace(nums, 0, n-2)

	withoutTakingZeroIndex := solveRob2ConstSpace(nums, 1, n-1)

	return max(withTakingZeroIndex, withoutTakingZeroIndex)
}

func solveRob2ConstSpace(nums []int, l int, r int) int {
	pp := nums[l]
	p := max(nums[l], nums[l+1])

	n := r - l + 1

	for k := 2; k < n; k++ {
		steal := nums[k+l] + pp
		notSteal := p

		tmp := max(steal, notSteal)
		pp = p
		p = tmp
	}
	return p
}

func (l *IntList) fillWith(v int) {
	for i := range *l {
		(*l)[i] = v
	}
}
