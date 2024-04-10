package top150

func Merge(nums1 []int, m int, nums2 []int, n int) {
	last := m + n - 1
	i, j := m-1, n-1
	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[last] = nums1[i]
			i--
		} else {
			nums1[last] = nums2[j]
			j--
		}
		last--
	}
	for j >= 0 {
		nums1[last] = nums2[j]
		j--
		last--
	}
}
