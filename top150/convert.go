package top150

import "strings"

func Convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	str := strings.Builder{}

	for i := 0; i < numRows; i++ {
		idx := i
		down := 2 * (numRows - 1 - i)
		up := 2 * i

		isDown := true
		for idx < len(s) {
			str.WriteByte(s[idx])

			if i == 0 {
				idx += down
			} else if i == numRows-1 {
				idx += up
			} else {
				if isDown {
					idx += down
				} else {
					idx += up
				}
				isDown = !isDown
			}

		}
	}

	return str.String()
}
