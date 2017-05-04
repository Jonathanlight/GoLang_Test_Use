package main

import "fmt"

func main(){
  const goldRadio float64 = 1.6212
  fmt.Println(goldRadio)

  //Constant Enumeration
  const(
    First = iota + 2
    Second
    Third
  )
  fmt.Println(First, Second, Third)
}
