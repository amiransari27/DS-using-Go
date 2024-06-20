package top150

func IsAnagram(s string, t string) bool {
   if len(s) != len(t) {
       return false
   }
   count := make([]int, 26)

   for i := 0; i < len(s); i++ {
       count[s[i]-'a'] += 1
   }

   for i := 0; i < len(t); i++ {
       count[t[i]-'a'] -= 1
   }

   for _, c := range count {
       if c > 0 {
           return false
       }
   }
   return true
}


