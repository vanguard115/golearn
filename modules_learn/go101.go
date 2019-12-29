package main

import "fmt"

func main() {
  fmt.Println("Hello Im here and Im here NOW!!")
  foo()
}

func foo() {
  fmt.Println("Im here in Foo now!")

  for i := 0; i < 100; i++ {
    if i%2 == 0 {
      fmt.Println(i)
    }
  }

}
