package dp

func IsMatch(s string, p string) bool {

	t := make([][]int, 21)
	for i := 0; i < 21; i++ {
		t[i] = make([]int, 21)
		for j := 0; j < 21; j++ {
			t[i][j] = -1
		}
	}
	return solve(0, 0, s, p, t)
}

func solve(i int, j int, s string, p string, t [][]int) bool {
	if j == len(p) {
		return i == len(s)
	}

	if t[i][j] != -1 {
		if t[i][j] == 0 {
			return false
		} else {
			return true
		}
	}

	firstCharMatch := false

	if i < len(s) && (s[i] == p[j] || p[j] == '.') {
		firstCharMatch = true
	}

	if j+1 < len(p) && p[j+1] == '*' {
		take := firstCharMatch && solve(i+1, j, s, p, t)
		notTake := solve(i, j+2, s, p, t)

		if take || notTake {
			t[i][j] = 1
		} else {
			t[i][j] = 0
		}

		return take || notTake
	} else {

		tmp := firstCharMatch && solve(i+1, j+1, s, p, t)

		if tmp {
			t[i][j] = 1
		} else {
			t[i][j] = 0
		}
		return tmp
	}
}
