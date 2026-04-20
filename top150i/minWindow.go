package top150i

import "math"

func minWindow(s string, t string) string {

	n := len(s)
	m := len(t)

	if m > n {
		return ""
	}

	mp := make(map[byte]int)
	for i := range t {
		mp[t[i]]++
	}

	reqCount := m
	i, j := 0, 0
	winSize := math.MaxInt

	start_i := 0

	//"ADOBECODEBANC"
	//"ABC"
	for j < n {
		ch := s[j]

		if mp[ch] > 0 {
			reqCount--
		}
		mp[ch]--

		for reqCount == 0 {

			//start shrinking the window
			currWinSize := j - i + 1
			if winSize > currWinSize {
				winSize = currWinSize
				start_i = i
			}
			mp[s[i]]++
			if mp[s[i]] > 0 {
				reqCount++
			}

			i++
		}

		j++
	}

	if winSize == math.MaxInt {
		return ""
	}

	return s[start_i : start_i+winSize]

}
