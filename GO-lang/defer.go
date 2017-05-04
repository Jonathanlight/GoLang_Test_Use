package main
import "fmt"

func main(){
  defer doSomething()
  defer doSomethingElseIf()
  doSomethingElse()

}
func doSomething(){
  fmt.Println("doSomething() 1er func ")
}

func doSomethingElse(){
  fmt.Println("doSomethingElse() 2eme func ")
}


func doSomethingElseIf(){
  fmt.Println("doSomethingElseIf() 3eme func ")
}
