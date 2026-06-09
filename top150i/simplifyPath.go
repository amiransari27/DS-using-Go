package top150i

import "strings"

func simplifyPath(path string) string {

	tokens := strings.Split(path, "/")
	stack := make([]string, 0)

	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "" || tokens[i] == "." {
			continue
		} else if tokens[i] == ".." {
			//pop the element
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, tokens[i])
		}

	}

	return "/" + strings.Join(stack, "/")
}
