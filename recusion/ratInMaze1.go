package recusion

// RatInMaze1 finds all valid paths from (0,0) to (n-1,n-1) in a maze using backtracking.
// Time Complexity: O(4^(n²)) in worst case - exponential number of paths
// Space Complexity: O(n²) for recursion stack depth + O(k*n²) for storing all k paths
func RatInMaze1(maze [][]int, n int) []string {

	result := []string{}

	solveRatInMaze1(0, 0, maze, []byte{}, &result, n)

	return result
}

func solveRatInMaze1(i int, j int, maze [][]int, path []byte, result *[]string, n int) {

	if i < 0 || i >= n || j < 0 || j >= n || maze[i][j] == 0 {
		return
	}

	if i == n-1 && j == n-1 {
		(*result) = append((*result), string(path))
		return
	}

	maze[i][j] = 0
	path = append(path, 'L')
	solveRatInMaze1(i, j-1, maze, path, result, n)
	path = path[:len(path)-1]

	path = append(path, 'R')
	solveRatInMaze1(i, j+1, maze, path, result, n)
	path = path[:len(path)-1]

	path = append(path, 'U')
	solveRatInMaze1(i-1, j, maze, path, result, n)
	path = path[:len(path)-1]

	path = append(path, 'D')
	solveRatInMaze1(i+1, j, maze, path, result, n)
	path = path[:len(path)-1]

	maze[i][j] = 1
}
