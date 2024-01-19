package leetcode

import "fmt"

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	plen := 0
	for {
		if plen < len(strs[0]) {
			ch := strs[0][plen]
			allWell := true

			for i := 1; i < len(strs); i++ {
				if len(strs[i]) > plen && strs[i][plen] == ch {
				} else {
					allWell = false
					break
				}
			}

			if allWell {
				plen++
			} else {
				break
			}

		} else {
			break
		}

	}

	fmt.Println(plen)
	ans := strs[0][:plen]

	return ans
}
