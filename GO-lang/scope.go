package main

import "fmt"

var iron int = 26

func main(){
  gold := 79
  fmt.Println(iron)
  fmt.Println(gold)
  anotherFunction()
}

func anotherFunction(){
  fmt.Println(iron)
}
