package top150

func Rotate(matrix [][]int) {

   // transform
   for i := 0; i < len(matrix); i++ {
       for j := i; j < len(matrix[0]); j++ {
           matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
       }
   }

   // reverse

   for i := 0; i < len(matrix); i++ {
       l, r := 0, len(matrix[0])-1
       for l < r {
           matrix[i][l], matrix[i][r] = matrix[i][r], matrix[i][l]

           l++
           r--
       }
   }

}


