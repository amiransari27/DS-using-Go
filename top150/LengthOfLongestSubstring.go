package top150

func LengthOfLongestSubstring(s string) int {
   n := len(s)
   l, r := 0, 0
   sl := 0

   charCount := make(map[byte]bool)
   for r < n {
       c := s[r]

       // release
       for charCount[c] {
           delete(charCount, s[l])
           l++
       }

       charCount[c] = true

       curLen := r - l + 1

       sl = max(sl, curLen)
       r++

   }
   return sl
}


