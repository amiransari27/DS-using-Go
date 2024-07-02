package top150

func ContainsNearbyDuplicate(nums []int, k int) bool {
   numSet := make(map[int]bool)

   i, j := 0, 0
   for j < len(nums) {
       // check i-j
       if j-i > k {
           // shrink
           delete(numSet, nums[i])
           i++
       }

       // check set
       if _, exist := numSet[nums[j]]; exist {
           return true
       } else {
           numSet[nums[j]] = true
       }
       j++
   }

   return false
}


