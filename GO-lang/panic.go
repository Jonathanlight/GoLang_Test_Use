package main

import (
  "io/ioutil"
  "fmt"
)

func main(){
  //read the whole file at once
  b, err := ioutil.ReadFile("system.txt")
  if err != nil {
    //is Exception
    panic("The program failed at file ")
  }
  for _, e := range b {
    fmt.Println(string(e))
  }
}
