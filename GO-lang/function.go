package main
import "fmt"

func main(){
  resultat := add(4,5,3)
  fmt.Println(resultat, doSomething())
}

func doSomething() string {
  text := "performi,g the doSomething"
  return text
}

func add(a int, b int, c int) int {
  rep := a + b + c
  return rep
}
