package top150

import (
   "slices"
)

func GroupAnagrams(strs []string) [][]string {

   groupMap := make(map[string][]string)

   for _, s := range strs {
       r := []rune(s)
       slices.Sort(r)

       t := string(r)
       if _, exist := groupMap[t]; exist {
           groupMap[t] = append(groupMap[t], s)
       } else {
           groupMap[t] = make([]string, 1)
           groupMap[t][0] = s
       }
   }

   ans := make([][]string, 0)
   for _, str_s := range groupMap {
       ans = append(ans, str_s)
   }
   return ans
}


