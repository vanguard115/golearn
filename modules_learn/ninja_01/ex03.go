package main

import (
  "fmt"
)

func main() {
  x := 42
  y := "James Bond"
  z := true

  text := fmt.Sprintf("%d %s %t", x, y ,z)

  fmt.Println(text)

}
