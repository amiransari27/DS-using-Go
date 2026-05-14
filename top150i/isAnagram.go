package top150i

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}
	charMap := [26]int{}

	for i := 0; i < len(s); i++ {
		charMap[s[i]-'a']++
	}

	for i := 0; i < len(t); i++ {
		charMap[t[i]-'a']--
	}

	for i := 0; i < len(charMap); i++ {
		if charMap[i] != 0 {
			return false
		}
	}

	return true
}
