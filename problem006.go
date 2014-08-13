package main

import (
  "fmt"
)

func main() {

  sumOfSquares := 0
  sums := 0

  for i:=1; i<=100; i++ {
    sumOfSquares += i * i
    sums += i 
  }

  fmt.Println(sums * sums - sumOfSquares)
}
