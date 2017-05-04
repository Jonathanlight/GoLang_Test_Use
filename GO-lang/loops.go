package main

import "fmt"

func main(){

  for i := 0; i <= 5; i++{
    fmt.Println(i)
  }

  fmt.Println(" ")
  fmt.Println(" ")

  myName := "Jonathan"
  for _, char := range myName {
    fmt.Println(string(char) + " ")
  }

  j := 0
  for j < 5 {
    fmt.Println("Golang")
    j++
    if j == 3{
      break
    }
  }

}
