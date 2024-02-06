package leetcode75

func FindDifference(nums1 []int, nums2 []int) [][]int {
	set1 := map[int]bool{}
	set2 := map[int]bool{}

	for _, v := range nums1 {
		set1[v] = true
	}

	for _, v := range nums2 {
		set2[v] = true
	}

	list1 := make([]int, 0)
	for key, _ := range set1 {
		if !set2[key] {
			list1 = append(list1, key)
		}
	}

	list2 := make([]int, 0)
	for key, _ := range set2 {
		if !set1[key] {
			list2 = append(list2, key)
		}
	}

	return [][]int{list1, list2}
}
