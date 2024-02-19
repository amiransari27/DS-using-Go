package leetcode75

type EPair struct {
	Char string
	Val  float64
}

func CalcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	adj := make(map[string][]EPair)

	for i, e := range equations {
		adj[e[0]] = append(adj[e[0]], EPair{Char: e[1], Val: values[i]})
		adj[e[1]] = append(adj[e[1]], EPair{Char: e[0], Val: 1 / values[i]})
	}

	res := make([]float64, 0)

	for _, q := range queries {
		src := q[0]
		dst := q[1]
		product := 1.0
		ans := -1.0

		if len(adj[src]) > 0 {
			visited := make(map[string]bool)
			dfs(src, dst, &visited, &adj, product, &ans)
		}

		res = append(res, ans)
	}
	return res
}

func dfs(src string, dst string, visited *map[string]bool, adj *map[string][]EPair, product float64, ans *float64) {
	if (*visited)[src] {
		return
	}

	(*visited)[src] = true

	if src == dst {
		*ans = product
		return
	}

	for _, p := range (*adj)[src] {
		dfs(p.Char, dst, visited, adj, product*p.Val, ans)
	}
}
