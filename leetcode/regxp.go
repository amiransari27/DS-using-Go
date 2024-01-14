package leetcode

func IsMatch(s string, p string) bool {

	var memo = make([][]*bool, len(s)+1)
	for i, _ := range memo {
		memo[i] = make([]*bool, len(p)+1)
	}
	res := match(0, 0, s, p, memo)
	return res
}

func match(si int, pi int, s string, p string, memo [][]*bool) bool {
	if len(p) == pi {
		return len(s) == si
	}

	if memo[si][pi] != nil {
		return *memo[si][pi]
	}

	fMatch := si < len(s) && (s[si] == p[pi] || p[pi] == "."[0])

	var ans bool
	if pi+1 < len(p) && p[pi+1] == "*"[0] {
		ans = (fMatch && match(si+1, pi, s, p, memo)) || match(si, pi+2, s, p, memo)
	} else {
		ans = fMatch && match(si+1, pi+1, s, p, memo)
	}
	memo[si][pi] = &ans
	return ans
}
