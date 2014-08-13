package main

import "fmt"

func main() {
  a, b, sum := 1,1,0
  for i:=3; a+b < 4000000; i++ {
    a, b = b, a+b
    if b % 2 == 0 {
      sum += b
    }
  }
  fmt.Println(sum)
}
