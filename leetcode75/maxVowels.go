package leetcode75

func MaxVowels(s string, k int) int {
	vowels := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	vc := 0
	ans := 0
	for i := 0; i < k; i++ {
		if vowels[s[i]] {
			vc++
		}
	}
	ans = vc

	for i := k; i < len(s); i++ {
		if vowels[s[i]] {
			vc++
		}
		if vowels[s[i-k]] {
			vc--
		}
		ans = max(ans, vc)
	}
	return ans
}
