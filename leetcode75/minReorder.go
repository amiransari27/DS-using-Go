package leetcode75

type Pair struct {
	City      int
	IsOrignal bool
}

// [[1,0],[1,2],[3,2],[3,4]]
func MinReorder(n int, connections [][]int) int {
	adj := make(map[int][]Pair)
	for _, conn := range connections {
		u := conn[0]
		v := conn[1]
		adj[u] = append(adj[u], Pair{City: v, IsOrignal: true})
		adj[v] = append(adj[v], Pair{City: u, IsOrignal: false})
	}
	visited := make(map[int]bool)
	ans := 0
	var dfs func(int)
	dfs = func(city int) {

		visited[city] = true

		for _, p := range adj[city] {
			c := p.City
			if !visited[c] {
				check := p.IsOrignal
				if check {
					ans++
				}
				dfs(c)
			}

		}
	}
	dfs(0)

	return ans
}
