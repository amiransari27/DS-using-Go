package top150i

func isValid(s string) bool {

	stack := make([]byte, 0)
	pMap := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else {
			if len(stack) > 0 {
				poped := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if pMap[poped] != s[i] {
					return false
				}
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
