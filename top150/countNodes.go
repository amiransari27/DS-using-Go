package top150

import "math"

func countNodes(root *TreeNode) int {
   if root == nil {
       return 0
   }

   lh := getLH(root)
   rh := getRH(root)

   if lh == rh {
       return int(math.Pow(2, float64(lh))) - 1
   }

   return countNodes(root.Left) + countNodes(root.Right) + 1
}

func getLH(root *TreeNode) int {
   tmp := root
   h := 0
   for tmp != nil {
       tmp = tmp.Left
       h++
   }
   return h
}

func getRH(root *TreeNode) int {
   tmp := root
   h := 0
   for tmp != nil {
       tmp = tmp.Right
       h++
   }
   return h
}

