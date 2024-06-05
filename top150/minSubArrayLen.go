package leetcodetop150

import "math"

func MinSubArrayLen(target int, nums []int) int {
   n := len(nums)
   sum := 0

   l, r := 0, 0
   minL := math.MaxInt

   for r < n {
       sum += nums[r]
       for sum >= target {
           minL = min(minL, r-l+1)

           sum -= nums[l]
           l++
       }

       r++
   }

   if minL == math.MaxInt {
       return 0
   } else {
       return minL
   }
}


