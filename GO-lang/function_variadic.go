package main
import "fmt"

func main(){
  r := add(600, 78, 95, 780)
  fmt.Println(r)
}

func add(nums ...int) int {
  var total int
  for _, n := range nums {
    total += n
  }
  return total
}
