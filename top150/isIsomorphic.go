package top150

func IsIsomorphic(s string, t string) bool {
   if len(s) != len(t) {
       return false
   }
   stMap := make(map[byte]byte)
   tsMap := make(map[byte]byte)

   for i := 0; i < len(s); i++ {

       ch1 := s[i]
       ch2 := t[i]

       _, existCh1 := stMap[ch1]
       _, existCh2 := tsMap[ch2]

       if (existCh1 && stMap[ch1] != ch2) || (existCh2 && tsMap[ch2] != ch1) {
           return false
       }

       stMap[ch1] = ch2
       tsMap[ch2] = ch1
   }

   return true
}




