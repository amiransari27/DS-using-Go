package problems

import "fmt"

func Toh(n int, rodA string, rodB string, rodC string) {
	if n == 0 {
		return
	}

	Toh(n-1, rodA, rodC, rodB)
	fmt.Printf("Move dist %d form %s to %s\n", n, rodA, rodB)
	Toh(n-1, rodC, rodB, rodA)
}
