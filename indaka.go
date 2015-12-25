package main

import (
  "fmt"
  s "strings"
)

/*
indak a
prefix ke-<isi>-ing akan dirubah menjadi prefix ka-<isi>-iang 
*/


func RubahKeBahasaMinangNotKa(kata string)string{
  if s.HasPrefix(kata, "ke"){
    // kering --> karing --> kariang
    // ke --> ka --> ka+n+<iang>
    kata := s.Replace(kata, "ke", "ka", 4)
    return s.Replace(kata, "ng", "ang", 100)
  }
  return s.Replace(kata, "ng", "ang" , 100) // set parameter to antisipate longer input
}

func test(kata string, expectation string)string{
  if RubahKeBahasaMinangNotKa(kata) == expectation{
    return "works"
  }
    return "not works"
}

func main(){
    // join in python "".join(list("golang"))
    w := "keriting"
    expectation := "karitiang"

    result := test(w, expectation)
    switch result{
    case "works":
        fmt.Println("hasil operasi:", result)
        fmt.Println("kata dalam bahasa Indonesia:", w)
        fmt.Println("kata dalam bahasa Minang:",RubahKeBahasaMinangNotKa(w))
    case "not works":
        fmt.Println("hasil operasi:", result)
        fmt.Println("kata dalam bahasa Indonesia:", w)
        fmt.Println("kata dalam bahasa Minang:",RubahKeBahasaMinangNotKa(w))
      }
    }



