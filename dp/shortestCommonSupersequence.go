package dp

func shortestCommonSupersequence(str1 string, str2 string) int {

	var solve func(s1 string, s2 string, i int, j int) int

	t := make(map[[2]int]int)

	solve = func(s1, s2 string, i, j int) int {

		if i == len(s1) {
			return len(s2) - j
		}

		if j == len(s2) {
			return len(s1) - i
		}

		key := [2]int{i, j}
		if val, exist := t[key]; exist {
			return val
		}

		var res int
		if s1[i] == s2[j] {
			res = 1 + solve(s1, s2, i+1, j+1)
		} else {

			res = 1 + min(solve(s1, s2, i+1, j), solve(s1, s2, i, j+1))
		}

		t[key] = res
		return res
	}

	return solve(str1, str2, 0, 0)
}

func shortestCommonSupersequenceBU(str1 string, str2 string) int {
	m := len(str1)
	n := len(str2)
	// use (m+1)x(n+1) table to handle empty prefixes
	t := make([][]int, m+1)
	for i := range t {
		t[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 || j == 0 {
				t[i][j] = i + j
			} else if str1[i-1] == str2[j-1] {
				t[i][j] = 1 + t[i-1][j-1]
			} else {
				t[i][j] = 1 + min(t[i-1][j], t[i][j-1])
			}
		}
	}

	return t[m][n]
}
