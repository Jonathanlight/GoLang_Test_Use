package main

import "fmt"

func main(){

  command := "create"

  switch command{
  case "create":
    fmt.Println("Create element")
  case "open":
    fmt.Println("Open element")
  case "close":
    fmt.Println("Close element")
  default:
    fmt.Println("Unrecognised command")
  }
}
