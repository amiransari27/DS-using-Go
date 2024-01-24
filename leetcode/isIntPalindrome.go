package leetcode

func IsIntPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var ans, num int = 0, x

	for num != 0 {
		ans = ans*10 + num%10
		num = num / 10
	}
	return ans == x
}
