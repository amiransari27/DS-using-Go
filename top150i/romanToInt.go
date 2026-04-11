package top150i

func romanToInt(s string) int {

	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	n := len(s)
	sum := 0

	for i := 0; i < n; i++ {
		ch := s[i]

		if i+1 < n && romanMap[ch] < romanMap[s[i+1]] {
			sum -= romanMap[ch]
		} else {
			sum += romanMap[ch]
		}
	}

	return sum

}
