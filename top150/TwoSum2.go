package top150

import (
   "sort"
)

type Pair struct {
   val int
   idx int
}

func TwoSum(nums []int, target int) []int {
   sl := make([]Pair, 0)

   for i, v := range nums {
       sl = append(sl, Pair{v, i})
   }

   sort.Slice(sl, func(i, j int) bool {
       return sl[i].val < sl[j].val
   })

   l, r := 0, len(nums)-1
   ans := make([]int, 2)
   for l < r {
       sum := sl[l].val + sl[r].val
       if sum == target {
           ans[0] = sl[l].idx
           ans[1] = sl[r].idx
           break
       } else if sum < target {
           l++
       } else {
           r--
       }
   }

   return ans
}


