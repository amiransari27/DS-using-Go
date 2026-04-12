package top150i

func longestCommonPrefix(strs []string) string {

	l := 0

	for {
		if len(strs[0]) > l {
			ch := strs[0][l]
			allWell := true

			for i := 1; i < len(strs); i++ {
				if l < len(strs[i]) && ch == strs[i][l] {
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

	return strs[0][:l]
}
