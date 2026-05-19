package top150i

func isHappy(n int) bool {

	cyMap := make(map[int]bool)

	for n != 1 {
		if _, exist := cyMap[n]; exist {
			return false
		} else {
			cyMap[n] = true
		}
		n = sumOfSquares(n)
	}

	return true
}

func sumOfSquares(n int) int {
	output := 0

	for n > 0 {
		digit := n % 10
		output += digit * digit
		n = n / 10
	}

	return output
}
