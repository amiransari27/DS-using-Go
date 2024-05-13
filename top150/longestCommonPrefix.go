package top150

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	l := 0

	for {
		if len(strs[0]) > l {
			char := strs[0][l]
			allWell := true
			for i := 1; i < len(strs); i++ {
				if l < len(strs[i]) && strs[i][l] == char {
					allWell = true
				} else {
					allWell = false
					break
				}
			}

			if allWell {
				l++
			} else {
				break
			}
		} else {
			break
		}
	}

	return strs[0][0:l]
}
