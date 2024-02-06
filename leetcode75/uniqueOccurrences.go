package leetcode75

func UniqueOccurrences(arr []int) bool {
	f := map[int]int{}
	for _, val := range arr {
		if f[val] != 0 {
			f[val] += 1
		} else {
			f[val] = 1
		}
	}
	s := map[int]bool{}
	for _, val := range f {
		if s[val] {
			return false
		}
		s[val] = true
	}
	return true
}
