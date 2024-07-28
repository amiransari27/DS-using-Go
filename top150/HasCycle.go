package top150

func HasCycle(head *ListNode) bool {
   slow, fast := head, head

   for fast.Next != nil && fast.Next.Next != nil {
       slow = slow.Next
       fast = fast.Next.Next

       if slow == fast {
           return true
       }
   }

   return false
}


