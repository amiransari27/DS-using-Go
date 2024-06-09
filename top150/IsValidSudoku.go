package top150

func IsValidSudoku(board [][]byte) bool {

   //check row
   for row := 0; row < 9; row++ {
       set := make(map[byte]bool)
       for col := 0; col < 9; col++ {
           if board[row][col] == '.' {
               continue
           }
           if set[board[row][col]] {
               return false
           }
           set[board[row][col]] = true
       }

   }

   //check column
   for col := 0; col < 9; col++ {
       set := make(map[byte]bool)
       for row := 0; row < 9; row++ {
           if board[row][col] == '.' {
               continue
           }
           if set[board[row][col]] {
               return false
           }
           set[board[row][col]] = true
       }
   }

   //check grid 3X3
   for sr := 0; sr < 9; sr += 3 {
       er := sr + 2
       for sc := 0; sc < 9; sc += 3 {
           ec := sc + 2

           if !validBox(board, sr, sc, er, ec) {
               return false
           }
       }
   }

   return true
}

func validBox(board [][]byte, sr int, sc int, er int, ec int) bool {
   set := make(map[byte]bool)
   for i := sr; i <= er; i++ {
       for j := sc; j <= ec; j++ {
           if board[i][j] == '.' {
               continue
           }
           if set[board[i][j]] {
               return false
           }
           set[board[i][j]] = true
       }
   }
   return true
}


