package top150

func isSymmetric(root *TreeNode) bool {
   if root == nil {
       return true
   }

   return check(root.Left, root.Right)
}

func check(l *TreeNode, r *TreeNode) bool {

   if l == nil && r == nil {
       return true
   }

   if l == nil || r == nil {
       return false
   }

   if l.Val == r.Val && check(l.Left, r.Right) && check(l.Right, r.Left) {
       return true
   }

   return false
}

