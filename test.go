/*
We have a two-dimensional board game involving snakes.  The board has two types of squares on it: +'s represent impassable squares where snakes cannot go, and 0's represent squares through which snakes can move.

Snakes may move in any of four directions - up, down, left, or right - one square at a time, but they will never return to a square that they've already visited.  If a snake enters the board on an edge square, we want to catch it at a different exit square on the board's edge.

The snake is familiar with the board and will take the route to the nearest reachable exit, in terms of the number of squares it has to move through to get there. Note that there may not be a reachable exit.

Here is an example board:

	col-->        0  1  2  3  4  5  6  7  8
	           +---------------------------
	row      0 |  +  +  +  +  +  +  +  0  0
	 |       1 |  +  +  0  0  0  0  0  +  +
	 |       2 |  0  0  0  0  0  +  +  0  +
	 v       3 |  +  +  0  +  +  +  +  0  0
	         4 |  +  +  0  0  0  0  0  0  +
	         5 |  +  +  0  +  +  0  +  0  +

Write a function that takes a rectangular board with only +'s and O's, along with a starting point on the edge of the board, and returns the coordinates of the nearest exit to which it can travel.  If there is a tie, return any of the nearest exits.
-----------------------------------------------------
Sample inputs:
board1 = [['+', '+', '+', '+', '+', '+', '+', '0', '0'],

	['+', '+', '0', '0', '0', '0', '0', '+', '+'],
	['0', '0', '0', '0', '0', '+', '+', '0', '+'],
	['+', '+', '0', '+', '+', '+', '+', '0', '0'],
	['+', '+', '0', '0', '0', '0', '0', '0', '+'],
	['+', '+', '0', '+', '+', '0', '+', '0', '+']]

start1_1 = (2, 0) # Expected output = (5, 2)
start1_2 = (0, 7) # Expected output = (0, 8)
start1_3 = (5, 2) # Expected output = (2, 0) or (5, 5)
start1_4 = (5, 5) # Expected output = (5, 7)

board2 = [['+', '+', '+', '+', '+', '+', '+'],

	['0', '0', '0', '0', '+', '0', '+'],
	['+', '0', '+', '0', '+', '0', '0'],
	['+', '0', '0', '0', '+', '+', '+'],
	['+', '+', '+', '+', '+', '+', '+']]

start2_1 = (1, 0) # Expected output = null (or a special value representing no possible exit)
start2_2 = (2, 6) # Expected output = null

board3 = [['+', '0', '+', '0', '+',],

	['0', '0', '+', '0', '0',],
	['+', '0', '+', '0', '+',],
	['0', '0', '+', '0', '0',],
	['+', '0', '+', '0', '+']]

start3_1 = (0, 1) # Expected output = (1, 0)
start3_2 = (4, 1) # Expected output = (3, 0)
start3_3 = (0, 3) # Expected output = (1, 4)
start3_4 = (4, 3) # Expected output = (3, 4)

board4 = [['+', '0', '+', '0', '+',],

	['0', '0', '0', '0', '0',],
	['+', '+', '+', '+', '+',],
	['0', '0', '0', '0', '0',],
	['+', '0', '+', '0', '+']]

start4_1 = (1, 0) # Expected output = (0, 1)
start4_2 = (1, 4) # Expected output = (0, 3)
start4_3 = (3, 0) # Expected output = (4, 1)
start4_4 = (3, 4) # Expected output = (4, 3)

board5 =  [['+', '0', '0', '0', '+',],

	['+', '0', '+', '0', '0',],
	['+', '0', '0', '0', '0',],
	['+', '0', '0', '0', '+']]

start5_1 = (0, 1) # Expected output = (0, 2)
start5_2 = (3, 1) # Expected output = (3, 2)
start5_3 = (1, 4) # Expected output = (2, 4)

board6 = [['+', '+', '+', '+', '+', '+', '+', '+'],

	['0', '0', '0', '0', '0', '0', '0', '0'],
	['+', '0', '0', '0', '0', '0', '0', '0'],
	['+', '0', '0', '0', '0', '0', '0', '+'],
	['0', '0', '0', '0', '0', '0', '0', '+'],
	['+', '+', '+', '+', '+', '+', '0', '+']]

start6_1 = (4,0) # Expected output = (1, 0)

All test cases:
findExit(board1, start1_1) => (5, 2)
findExit(board1, start1_2) => (0, 8)
findExit(board1, start1_3) => (2, 0) or (5, 5)
findExit(board1, start1_4) => (5, 7)
findExit(board2, start2_1) => null (or a special value representing no possible exit)
findExit(board2, start2_2) => null
findExit(board3, start3_1) => (1, 0)
findExit(board3, start3_2) => (3, 0)
findExit(board3, start3_3) => (1, 4)
findExit(board3, start3_4) => (3, 4)
findExit(board4, start4_1) => (0, 1)
findExit(board4, start4_2) => (0, 3)
findExit(board4, start4_3) => (4, 1)
findExit(board4, start4_4) => (4, 3)
findExit(board5, start5_1) => (0, 2)
findExit(board5, start5_2) => (3, 2)
findExit(board5, start5_3) => (2, 4)
findExit(board6, start6_1) => (1, 0)

Complexity Analysis:

r: number of rows in the board
c: number of columns in the board
*/
package main

import "fmt"

func testmain() {
	// board1 := [][]byte{
	//   []byte{'+', '+', '+', '+', '+', '+', '+', '0', '0'},
	//   []byte{'+', '+', '0', '0', '0', '0', '0', '+', '+'},
	//   []byte{'0', '0', '0', '0', '0', '+', '+', '0', '+'},
	//   []byte{'+', '+', '0', '+', '+', '+', '+', '0', '0'},
	//   []byte{'+', '+', '0', '0', '0', '0', '0', '0', '+'},
	//   []byte{'+', '+', '0', '+', '+', '0', '+', '0', '+'},
	// }
	// start1_1 := []int{2, 0} // Expected output = {5, 2}
	// start1_2 := []int{0, 7} // Expected output = {0, 8}
	// start1_3 := []int{5, 2} // Expected output = {2, 0} or {5, 5}
	// start1_4 := []int{5, 5} // Expected output = {5, 7}

	// board2 := [][]byte{
	//     []byte{'+', '+', '+', '+', '+', '+', '+'},
	//     []byte{'0', '0', '0', '0', '+', '0', '+'},
	//     []byte{'+', '0', '+', '0', '+', '0', '0'},
	//     []byte{'+', '0', '0', '0', '+', '+', '+'},
	//     []byte{'+', '+', '+', '+', '+', '+', '+'},
	// }
	// start2_1 := []int{1, 0} // Expected output = null {or a special value representing no possible exit}
	// start2_2 := []int{2, 6} // Expected output = null

	// board3 := [][]byte{
	//   []byte{'+', '0', '+', '0', '+'},
	//   []byte{'0', '0', '+', '0', '0'},
	//   []byte{'+', '0', '+', '0', '+'},
	//   []byte{'0', '0', '+', '0', '0'},
	//   []byte{'+', '0', '+', '0', '+'},
	// }
	// start3_1 := []int{0, 1} // Expected output = {1, 0}
	// start3_2 := []int{4, 1} // Expected output = {3, 0}
	// start3_3 := []int{0, 3} // Expected output = {1, 4}
	// start3_4 := []int{4, 3} // Expected output = {3, 4}

	// board4 := [][]byte{
	//   []byte{'+', '0', '+', '0', '+'},
	//   []byte{'0', '0', '0', '0', '0'},
	//   []byte{'+', '+', '+', '+', '+'},
	//   []byte{'0', '0', '0', '0', '0'},
	//   []byte{'+', '0', '+', '0', '+'},
	// }
	// start4_1 := []int{1, 0} // Expected output = {0, 1}
	// start4_2 := []int{1, 4} // Expected output = {0, 3}
	// start4_3 := []int{3, 0} // Expected output = {4, 1}
	// start4_4 := []int{3, 4} // Expected output = {4, 3}

	// board5 := [][]byte{
	//   []byte{'+', '0', '0', '0', '+'},
	//   []byte{'+', '0', '+', '0', '0'},
	//   []byte{'+', '0', '0', '0', '0'},
	//   []byte{'+', '0', '0', '0', '+'},
	// }
	// start5_1 := []int{0, 1} // Expected output = {0, 2}
	// start5_2 := []int{3, 1} // Expected output = {3, 2}
	// start5_3 := []int{1, 4} // Expected output = {2, 4}

	board6 := [][]byte{
		[]byte{'+', '+', '+', '+', '+', '+', '+', '+'},
		[]byte{'0', '0', '0', '0', '0', '0', '0', '0'},
		[]byte{'+', '0', '0', '0', '0', '0', '0', '0'},
		[]byte{'+', '0', '0', '0', '0', '0', '0', '+'},
		[]byte{'0', '0', '0', '0', '0', '0', '0', '+'},
		[]byte{'+', '+', '+', '+', '+', '+', '0', '+'},
	}
	start6_1 := []int{4, 0} // Expected output = {1, 0}

	result := findExit(board6, start6_1)
	fmt.Println("Exit found at:", result)

}

func findExit(board [][]byte, start []int) []int {
	r := len(board)
	c := len(board[0])

	// BFS to find shortest path
	visited := make([][]bool, r)
	for i := range visited {
		visited[i] = make([]bool, c)
	}

	queue := [][]int{start}
	visited[start[0]][start[1]] = true

	// Four directions: up, down, left, right
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// Explore all 4 neighbors
		for _, dir := range directions {
			nr, nc := curr[0]+dir[0], curr[1]+dir[1]

			// Check bounds
			if nr < 0 || nr >= r || nc < 0 || nc >= c {
				continue
			}

			// Check if passable and unvisited
			if board[nr][nc] == '+' || visited[nr][nc] {
				continue
			}

			// Check if reached an edge (and not the starting position)
			if (nr == 0 || nr == r-1 || nc == 0 || nc == c-1) &&
				!(nr == start[0] && nc == start[1]) {
				return []int{nr, nc}
			}

			visited[nr][nc] = true
			queue = append(queue, []int{nr, nc})
		}
	}

	// No exit found
	return []int{-1, -1}
}

func findPassableLanes(board [][]byte) [][]int {

	r := len(board)
	c := len(board[0])

	// check rows
	rRes := make([]int, 0)
	for i := 0; i < r; i++ {
		isPassThrough := true
		for j := 0; j < c; j++ {
			if board[i][j] == '+' {
				isPassThrough = false
				break
			}

		}
		if isPassThrough == true {
			rRes = append(rRes, i)
		}
	}

	// check col
	cRes := make([]int, 0)
	for j := 0; j < c; j++ {
		isPassThrough := true
		for i := 0; i < r; i++ {
			if board[i][j] == '+' {
				isPassThrough = false
				break
			}
		}
		if isPassThrough == true {
			cRes = append(cRes, j)
		}
	}

	return [][]int{rRes, cRes}
}

/*
package main
import "fmt"

func main() {
  // board1 := [][]byte{
  //   []byte{'+', '+', '+', '+', '+', '+', '+', '0', '0'},
  //   []byte{'+', '+', '0', '0', '0', '0', '0', '+', '+'},
  //   []byte{'0', '0', '0', '0', '0', '+', '+', '0', '+'},
  //   []byte{'+', '+', '0', '+', '+', '+', '+', '0', '0'},
  //   []byte{'+', '+', '0', '0', '0', '0', '0', '0', '+'},
  //   []byte{'+', '+', '0', '+', '+', '0', '+', '0', '+'},
  // }
  // start1_1 := []int{2, 0} // Expected output = {5, 2}
  // start1_2 := []int{0, 7} // Expected output = {0, 8}
  // start1_3 := []int{5, 2} // Expected output = {2, 0} or {5, 5}
  // start1_4 := []int{5, 5} // Expected output = {5, 7}

  // board2 := [][]byte{
  //     []byte{'+', '+', '+', '+', '+', '+', '+'},
  //     []byte{'0', '0', '0', '0', '+', '0', '+'},
  //     []byte{'+', '0', '+', '0', '+', '0', '0'},
  //     []byte{'+', '0', '0', '0', '+', '+', '+'},
  //     []byte{'+', '+', '+', '+', '+', '+', '+'},
  // }
  // start2_1 := []int{1, 0} // Expected output = null {or a special value representing no possible exit}
  // start2_2 := []int{2, 6} // Expected output = null

  // board3 := [][]byte{
  //   []byte{'+', '0', '+', '0', '+'},
  //   []byte{'0', '0', '+', '0', '0'},
  //   []byte{'+', '0', '+', '0', '+'},
  //   []byte{'0', '0', '+', '0', '0'},
  //   []byte{'+', '0', '+', '0', '+'},
  // }
  // start3_1 := []int{0, 1} // Expected output = {1, 0}
  // start3_2 := []int{4, 1} // Expected output = {3, 0}
  // start3_3 := []int{0, 3} // Expected output = {1, 4}
  // start3_4 := []int{4, 3} // Expected output = {3, 4}

  // board4 := [][]byte{
  //   []byte{'+', '0', '+', '0', '+'},
  //   []byte{'0', '0', '0', '0', '0'},
  //   []byte{'+', '+', '+', '+', '+'},
  //   []byte{'0', '0', '0', '0', '0'},
  //   []byte{'+', '0', '+', '0', '+'},
  // }
  // start4_1 := []int{1, 0} // Expected output = {0, 1}
  // start4_2 := []int{1, 4} // Expected output = {0, 3}
  // start4_3 := []int{3, 0} // Expected output = {4, 1}
  // start4_4 := []int{3, 4} // Expected output = {4, 3}

  // board5 := [][]byte{
  //   []byte{'+', '0', '0', '0', '+'},
  //   []byte{'+', '0', '+', '0', '0'},
  //   []byte{'+', '0', '0', '0', '0'},
  //   []byte{'+', '0', '0', '0', '+'},
  // }
  // start5_1 := []int{0, 1} // Expected output = {0, 2}
  // start5_2 := []int{3, 1} // Expected output = {3, 2}
  // start5_3 := []int{1, 4} // Expected output = {2, 4}

  board6 := [][]byte{
      []byte{'+', '+', '+', '+', '+', '+', '+', '+'},
      []byte{'0', '0', '0', '0', '0', '0', '0', '0'},
      []byte{'+', '0', '0', '0', '0', '0', '0', '0'},
      []byte{'+', '0', '0', '0', '0', '0', '0', '+'},
      []byte{'0', '0', '0', '0', '0', '0', '0', '+'},
      []byte{'+', '+', '+', '+', '+', '+', '0', '+'},
  }
  start6_1 := []int{4, 0} // Expected output = {1, 0}

  findExit(board6, start6_1)

}

func findExit(board [][]byte, start []int){

  r := len(board)
  c := len(board[0])
  pathMemo := make([][]int, r)
  for i, _ := range pathMemo {
    pathMemo[i] = make([]int, c)
  }
  ansMap := make(map[int][]int)
  pathCounter := 0

  findHelper(board, pathMemo, ansMap, start, &pathCounter, r, c, true)

  fmt.Println(ansMap)

}

func findHelper(board[][]byte, pathMemo[][]int, ansMap map[int][]int, curr []int, pathCounter *int, r int, c int, isStart bool){

    fmt.Println("till 1", curr)
    if curr[0] < 0 || curr[0] >r || curr[1] < 0 || curr[1] > c{
      return
    }
    if !isStart && (curr[0] == r || curr[1]== c){
      fmt.Println(curr, *pathCounter)
        ansMap[*pathCounter] = []int{curr[0] , curr[1]}
        return
    }
    *pathCounter += 1

    if pathMemo[curr[0]][curr[1]] == '0' && pathMemo[curr[0]][curr[1]] == 0 {
      // check top

      pathMemo[curr[0]][curr[1]] = 1
      findHelper(board, pathMemo, ansMap, []int{curr[0] -1, curr[1] }, pathCounter, r, c, false)
      pathMemo[curr[0]][curr[1]] = 0
      // check bottom
      pathMemo[curr[0]][curr[1]] = 1
      findHelper(board, pathMemo, ansMap, []int{curr[0] +1, curr[1] }, pathCounter, r, c, false)
      pathMemo[curr[0]][curr[1]] = 0
      // check left
      pathMemo[curr[0]][curr[1]] = 1
      findHelper(board, pathMemo, ansMap, []int{curr[0], curr[1] -1 }, pathCounter, r, c, false)
      pathMemo[curr[0]][curr[1]] = 0
      // check right
      pathMemo[curr[0]][curr[1]] = 1
      findHelper(board, pathMemo, ansMap, []int{curr[0], curr[1] +1 }, pathCounter, r, c, false)
      pathMemo[curr[0]][curr[1]] = 0
    }



}

func findPassableLanes(board [][]byte) [][]int{

    r := len(board)
    c := len(board[0])

    // check rows
    rRes := make([]int,0)
    for i:= 0; i< r; i++{
      isPassThrough := true
       for j:= 0; j< c ; j++ {
          if board[i][j] == '+'{
            isPassThrough = false
            break
          }

       }
       if isPassThrough == true {
         rRes = append(rRes, i)
       }
    }

    // check col
    cRes := make([]int,0)
    for j:= 0; j< c; j++{
      isPassThrough := true
       for i:= 0; i< r ; i++ {
          if board[i][j] == '+'{
            isPassThrough = false
            break
          }
       }
       if isPassThrough == true {
         cRes = append(cRes, j)
       }
    }


    return [][]int{rRes, cRes}
}

*/
