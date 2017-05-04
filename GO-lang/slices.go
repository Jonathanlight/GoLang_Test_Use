package main
import "fmt"

func main() {
  s := []string{}
  r := make([]string, 3)

  fmt.Println(s)
  fmt.Println(r)

  r[0] = "hello"
  r[1] = "hi"
  r[2] = "whattup"

  fmt.Println( r[0], r[1], r[2] )
}
