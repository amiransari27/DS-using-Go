package top150

func GameOfLife(board [][]int) {

   m := len(board)
   n := len(board[0])

   dir := [][]int{
       {0, -1}, {-1, 0}, {0, 1}, {1, 0},
       {-1, -1}, {-1, 1}, {1, 1}, {1, -1},
   }

   for i := 0; i < m; i++ {
       for j := 0; j < n; j++ {

           ln := 0
           for _, v := range dir {
               ni := v[0] + i
               nj := v[1] + j
               if ni >= 0 && ni < m && nj >= 0 && nj < n {
                   if board[ni][nj] == 1 || board[ni][nj] == -1 {
                       ln++
                   }
               }
           }

           if board[i][j] == 1 && (ln < 2 || ln > 3) {
               board[i][j] = -1
           }

           if board[i][j] == 0 && ln == 3 {
               board[i][j] = 2
           }

       }
   }

   for i := 0; i < m; i++ {
       for j := 0; j < n; j++ {
           if board[i][j] >= 1 {
               board[i][j] = 1
           } else {
               board[i][j] = 0
           }
       }
   }
}


