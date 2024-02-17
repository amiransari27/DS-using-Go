package leetcode75

func FindCircleNum(isConnected [][]int) int {
	visited := make(map[int]bool)

	// make graph
	adj := make(map[int][]int)
	for i := range isConnected {
		for j := range isConnected[i] {
			if isConnected[i][j] == 1 {
				adj[i] = append(adj[i], j)
				adj[j] = append(adj[j], i)
			}
		}
	}

	province := 0
	for c := range adj {
		if !visited[c] {
			dfs_findCircleNum(&adj, c, &visited)
			province++
		}
	}

	return province
}

func dfs_findCircleNum(adj *map[int][]int, i int, visited *map[int]bool) {
	(*visited)[i] = true

	for _, city := range (*adj)[i] {
		if !(*visited)[city] {
			dfs_findCircleNum(adj, city, visited)
		}
	}
}
