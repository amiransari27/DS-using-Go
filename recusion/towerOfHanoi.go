package recusion

import "fmt"

func TowerOfHanoi(n int, from string, to string, aux string) int {

	if n == 1 {
		fmt.Printf("Move disk %d from rod %s to rod %s\n", n, from, to)
		return 1
	}
	count := 0

	count = TowerOfHanoi(n-1, from, aux, to)
	fmt.Printf("Move disk %d from rod %s to rod %s\n", n, from, to)
	count += 1
	count += TowerOfHanoi(n-1, aux, to, from)

	return count
}
