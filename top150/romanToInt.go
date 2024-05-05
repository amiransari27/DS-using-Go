package top150

func RomanToInt(s string) int {
	num := 0
	n := len(s)

	for i, char := range s {

		if char == 'M' {
			num += 1000
		} else if char == 'D' {
			num += 500
		} else if char == 'C' && i+1 < n && (s[i+1] == 'M' || s[i+1] == 'D') {
			num -= 100
		} else if char == 'C' {
			num += 100
		} else if char == 'L' {
			num += 50
		} else if char == 'X' && i+1 < n && (s[i+1] == 'L' || s[i+1] == 'C') {
			num -= 10
		} else if char == 'X' {
			num += 10
		} else if char == 'V' {
			num += 5
		} else if char == 'I' && i+1 < n && (s[i+1] == 'X' || s[i+1] == 'V') {
			num -= 1
		} else if char == 'I' {
			num += 1
		}
	}

	return num
}
