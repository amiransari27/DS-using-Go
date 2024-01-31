package leetcode75

import "strings"

func MergeAlternately(word1 string, word2 string) string {

	var p1, p2 int = 0, 0
	str := strings.Builder{}
	turn := true

	for p1 < len(word1) && p2 < len(word2) {
		if turn {
			str.WriteByte(word1[p1])
			p1++
		} else {
			str.WriteByte(word2[p2])
			p2++
		}
		turn = !turn
	}

	if p1 <= len(word1)-1 {
		str.WriteString(word1[p1:])
	}
	if p2 <= len(word2)-1 {
		str.WriteString(word2[p2:])
	}

	return str.String()
}
