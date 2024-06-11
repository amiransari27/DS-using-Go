package top150

func SpiralOrder(matrix [][]int) []int {
   m, n := len(matrix), len(matrix[0])
   // left to right = 0
   // top to bottom = 1
   // right to left = 2
   // bottom to top = 3
   dir := 0

   result := make([]int, 0)

   l, r, t, b := 0, n-1, 0, m-1

   for l <= r && t <= b {
       if dir == 0 {
           // left to right
           // constant row - t

           for i := l; i <= r; i++ {
               result = append(result, matrix[t][i])
           }
           t++
       }

       if dir == 1 {
           // top to bottom
           // constant col - r

           for i := t; i <= b; i++ {
               result = append(result, matrix[i][r])
           }
           r--
       }

       if dir == 2 {
           // right to left
           // constant row - b

           for i := r; i >= l; i-- {
               result = append(result, matrix[b][i])
           }
           b--
       }

       if dir == 3 {
           // bottom to top
           // constant col - l

           for i := b; i >= t; i-- {
               result = append(result, matrix[i][l])
           }
           l++
       }

       dir++

       if dir == 4 {
           dir = 0
       }

   }

   return result
}


