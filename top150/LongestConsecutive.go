package top150

func LongestConsecutive(nums []int) int {
   numSet := make(map[int]bool)

   // Populate the set for fast lookup
   for _, num := range nums {
       numSet[num] = true
   }

   maxLength := 0

   // Traverse through each number and find the consecutive sequences
   for _, num := range nums {
       // Check if num is the start of a sequence
       if !numSet[num-1] {
           currentNum := num
           currentLength := 1

           // Extend the sequence to the right
           for numSet[currentNum+1] {
               currentNum++
               currentLength++
           }

           // Update the maximum length found
           if currentLength > maxLength {
               maxLength = currentLength
           }
       }
   }

   return maxLength
}

