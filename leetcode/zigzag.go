package leetcode

func Convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	result := []byte{}

	for i := 0; i < numRows; i++ {
		idx := i
		driveDown := 2 * (numRows - 1 - i)
		driveUp := 2 * i
		driveToggle := true

		for idx < len(s) {
			result = append(result, s[idx])

			if i == 0 {
				idx += driveDown
			} else if i == numRows-1 {
				idx += driveUp
			} else {
				if driveToggle {
					idx += driveDown
				} else {
					idx += driveUp
				}
				driveToggle = !driveToggle
			}

		}
	}
	return string(result)

}
