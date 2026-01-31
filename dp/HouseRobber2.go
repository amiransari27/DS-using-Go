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

func (l *IntList) fillWith(v int) {
	for i := range *l {
		(*l)[i] = v
	}
}
