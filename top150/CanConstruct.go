package top150

func CanConstruct(ransomNote string, magazine string) bool {
   mMap := make(map[byte]int)

   for i := 0; i < len(magazine); i++ {
       mMap[magazine[i]] += 1
   }

   for i := 0; i < len(ransomNote); i++ {
       if mMap[ransomNote[i]] == 0 {
           return false
       }
       mMap[ransomNote[i]] -= 1
   }
   return true
}


