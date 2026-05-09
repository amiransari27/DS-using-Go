package top150i

func canConstruct(ransomNote string, magazine string) bool {

	count := [26]int{}

	for _, ch := range magazine {
		count[ch-'a'] += 1
	}

	for _, ch := range ransomNote {
		count[ch-'a']--

		if count[ch-'a'] < 0 {
			return false
		}
	}

	return true
}
