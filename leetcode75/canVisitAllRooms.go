package leetcode75

func CanVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))

	dfs_visit_room(&rooms, 0, &visited)
	for _, visit := range visited {
		if !visit {
			return false
		}
	}
	return true
}

func dfs_visit_room(room *[][]int, source int, visited *[]bool) {

	(*visited)[source] = true

	for _, key := range (*room)[source] {
		if !(*visited)[key] {
			dfs_visit_room(room, key, visited)
		}
	}
}
