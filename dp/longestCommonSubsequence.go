package dp

func longestCommonSubsequence(text1 string, text2 string) int {

	t := make(map[[2]int]int)
	return solveLongestCommonSubsequence(text1, text2, 0, 0, &t)
}

func solveLongestCommonSubsequence(text1 string, text2 string, i int, j int, t *map[[2]int]int) int {

	if i >= len(text1) || j >= len(text2) {
		return 0
	}
	key := [2]int{i, j}
	if v, exist := (*t)[key]; exist {
		return v
	}

	if text1[i] == text2[j] {
		(*t)[key] = 1 + solveLongestCommonSubsequence(text1, text2, i+1, j+1, t)
		return (*t)[key]
	} else {
		take_text1 := solveLongestCommonSubsequence(text1, text2, i, j+1, t)

		take_text2 := solveLongestCommonSubsequence(text1, text2, i+1, j, t)

		(*t)[key] = max(take_text1, take_text2)
		return (*t)[key]
	}

}

func longestCommonSubsequenceBU(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)

	if m == 0 || n == 0 {
		return 0
	}

	t := make([][]int, m+1)
	for i := range t {
		t[i] = make([]int, n+1)
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {

			if text1[i-1] == text2[j-1] {
				t[i][j] = 1 + t[i-1][j-1]
			} else {
				t[i][j] = max(t[i-1][j], t[i][j-1])
			}
		}
	}

	return t[m][n]

}

func longestCommonSubsequenceStringBU(text1 string, text2 string) string {
	m := len(text1)
	n := len(text2)

	if m == 0 || n == 0 {
		return ""
	}

	t := make([][]int, m+1)
	for i := range t {
		t[i] = make([]int, n+1)
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {

			if text1[i-1] == text2[j-1] {
				t[i][j] = 1 + t[i-1][j-1]
			} else {
				t[i][j] = max(t[i-1][j], t[i][j-1])
			}
		}
	}

	str := make([]byte, t[m][n])

	i := m
	j := n

	for i > 0 && j > 0 {
		if text1[i-1] == text2[j-1] {
			str[t[i][j]-1] = text1[i-1]
			i--
			j--
		} else {
			if t[i-1][j] > t[i][j-1] {
				i--
			} else {
				j--
			}
		}

	}

	return string(str)

}
