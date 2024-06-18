package top150

import (
   "strings"
)

func WordPattern(pattern string, s string) bool {

   words := strings.Split(s, " ")
   if len(pattern) != len(words) {
       return false
   }

   strToChar := make(map[string]byte)
   charToStr := make(map[byte]bool)

   for i, word := range words {
       if _, exist := strToChar[word]; exist {
           if pattern[i] != strToChar[word] {
               return false
           }
       } else {
           if charToStr[pattern[i]] {
               return false
           }
           strToChar[word] = pattern[i]
           charToStr[pattern[i]] = true
       }
   }
   return true
}


