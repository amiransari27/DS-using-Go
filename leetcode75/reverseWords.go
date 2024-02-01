package leetcode75

import (
	"strings"
)

func ReverseWords(s string) string {

	sArr := strings.Split(strings.Trim(s, " "), " ")
	str := strings.Builder{}
	for i := len(sArr) - 1; i >= 0; i-- {
		if sArr[i] != "" {
			str.WriteString(sArr[i])
			str.WriteString(" ")
		}
	}

	return strings.Trim(str.String(), " ")
}
