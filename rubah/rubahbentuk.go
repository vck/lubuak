package main

import (
	f "fmt"
	s "strings"
)

func SufixIang(kata string)string{
  if s.HasPrefix(kata, "ke"){
    // kering --> karing --> kariang
    // ke --> ka --> ka+n+<iang>
    kata := s.Replace(kata, "ke", "ka", 4)
    return s.Replace(kata, "ng", "ang", 100)
  }
  return s.Replace(kata, "ng", "ang" , 100) // set parameter to antisipate longer input
}

func SufixMinang(kata string)string{
  sz := len(kata)
  if sz > 0 && kata[sz-1] == 'a' {
    kata = kata[:sz-1]+"o"
    return kata
  }
  return kata
}

func SufixIndonesia(kata string)string{
  sz := len(kata)
  if sz > 0 && kata[sz-1] == 'o' {
    kata = kata[:sz-1]+"a"
    return kata
  }
  return kata
}


func main(){
	f.Println(SufixMinang("luka"))
	f.Println(SufixIndonesia("duo"))
}
