package recusion

func MergeSort(nums []int, l int, r int) {

	if l == r {
		return
	}

	mid := l + (r-l)/2

	MergeSort(nums, l, mid)
	MergeSort(nums, mid+1, r)

	merge(nums, l, mid, r)
}

func merge(nums []int, l int, m int, r int) {

	n1 := m - l + 1
	n2 := r - m
	k := l

	L := make([]int, n1)
	R := make([]int, n2)

	for i := 0; i < n1; i++ {
		L[i] = nums[k]
		k++
	}
	for i := 0; i < n2; i++ {
		R[i] = nums[k]
		k++
	}

	// merge
	i := 0
	j := 0
	k = l
	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			nums[k] = L[i]
			i++
		} else {
			nums[k] = R[j]
			j++
		}
		k++
	}

	for i < n1 {
		nums[k] = L[i]
		i++
		k++
	}
	for j < n2 {
		nums[k] = R[j]
		j++
		k++
	}

}
