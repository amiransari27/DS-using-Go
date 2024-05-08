package top150

import "strings"

func LengthOfLastWord(s string) int {

	s = strings.Trim(s, " ")
	strArr := strings.Split(s, " ")
	lastStr := strArr[len(strArr)-1]

	return len(lastStr)
}

func LengthOfLastWord2(s string) int {
	strArr := strings.Fields(s)

	return len(strArr[len(strArr)-1])
}
