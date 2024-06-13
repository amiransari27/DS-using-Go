package top150

func SetZeroes(matrix [][]int) {

   m, n := len(matrix), len(matrix[0])

   rows := make([]int, m)
   cols := make([]int, n)

   for idx := range rows {
       rows[idx] = 1
   }

   for idx := range cols {
       cols[idx] = 1
   }

   for i := 0; i < m; i++ {
       for j := 0; j < n; j++ {
           if matrix[i][j] == 0 {
               rows[i] = 0
               cols[j] = 0
           }
       }
   }

   for idx := range rows {
       if rows[idx] == 0 {
           for i := 0; i < n; i++ {
               matrix[idx][i] = 0
           }
       }
   }

   for idx := range cols {
       if cols[idx] == 0 {
           for i := 0; i < m; i++ {
               matrix[i][idx] = 0
           }
       }
   }
}


