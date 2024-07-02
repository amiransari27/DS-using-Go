package top150

func IsHappy(n int) bool {

   visit := make(map[int]bool)

   for !visit[n] {
       visit[n] = true

       n = sumOfSquares(n)

       if n == 1 {
           return true
       }
   }
   return false

}

func sumOfSquares(n int) int {
   output := 0

   for n > 0 {

       digit := n % 10
       output += digit * digit
       n = n / 10
   }

   return output
}


