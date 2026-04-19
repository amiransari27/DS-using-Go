package top150i

func lengthOfLongestSubstring(s string) int {

	res := 0

	chMap := make(map[byte]bool)

	l, r := 0, 0

	for ; r < len(s); r++ {

		ch := s[r]

		for chMap[ch] == true {
			chMap[s[l]] = false
			l++
		}

		chMap[ch] = true

		res = max(res, r-l+1)
	}

	return res
}
