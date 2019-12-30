package main

import (
  "fmt"
)

type starShip int
var x starShip
var y int

func main() {

  fmt.Println(x)
  x  = 42
  fmt.Printf("%T\n",x )
  fmt.Println(x)

  y = int(x)
  fmt.Printf("%T\t%d",y,y)
}
