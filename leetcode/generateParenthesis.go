package leetcode

import "strings"

func GenerateParenthesis(n int) []string {
	list := make([][]string, n+1)

	list[0] = []string{""}
	list[1] = []string{"()"}

	for i := 2; i <= n; i++ {
		left := 0
		right := i - 1
		for right >= 0 && left < i {
			for _, s1 := range list[left] {
				for _, s2 := range list[right] {
					str := strings.Builder{}
					str.WriteString("(")
					str.WriteString(s1)
					str.WriteString(")")
					str.WriteString(s2)
					list[i] = append(list[i], str.String())
				}
			}

			left++
			right--
		}

	}
	return list[n]
}

func GenerateParenthesisRec(n int) []string {
	res := make([]string, 0)
	dfs_GenerateParenthesisRec(&res, n, 0, 0, "")
	return res
}
func dfs_GenerateParenthesisRec(res *[]string, n int, open int, close int, str string) {
	if open != n {
		dfs_GenerateParenthesisRec(res, n, open+1, close, str+"(")
	}

	if close != n && close < open {
		dfs_GenerateParenthesisRec(res, n, open, close+1, str+")")
	}

	if open == n && close == n {
		*res = append(*res, str)
	}
}
