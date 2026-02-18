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
