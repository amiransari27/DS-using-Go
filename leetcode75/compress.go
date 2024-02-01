package leetcode75

import (
	"strconv"
	"strings"
)

func compress(chars []byte) int {
	var cChar byte
	var cCount int

	str := strings.Builder{}

	for _, c := range chars {
		if c == cChar {
			cCount++
		} else {
			cChar = c
			if cCount > 1 {
				str.WriteString(strconv.Itoa(cCount))
			}
			str.WriteByte(cChar)
			cCount = 1
		}
	}

	if cCount > 1 {
		str.WriteString(strconv.Itoa(cCount))
	}

	tmpStr := str.String()

	for i := 0; i < len(tmpStr); i++ {
		chars[i] = tmpStr[i]
	}

	return str.Len()

}
