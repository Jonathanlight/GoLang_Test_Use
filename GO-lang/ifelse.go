package main

import "fmt"

func main(){

  nb1, nb2 := 5, 2
  if nb1 > nb2 {
    fmt.Println(nb1 , " is indeed greater than " , nb2)
  } else {
    fmt.Println(nb1 , " is not greater than " , nb2)
  }

  if nb1 >= 2 && nb2 <= 2 {
    fmt.Println(nb1 , " in intervale " , nb2)
  } else {
    fmt.Println(nb1 , " is not in intervale " , nb2)
  }

}
