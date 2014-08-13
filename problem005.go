package main

import (
  "fmt"
)

func main() {
  multiples := []int{2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
  finalMulti := []int{1}
  for len(multiples) > 0 {
    head := multiples[0]
    tail := multiples[1:]
    for i := 0; i < len(tail); i++ {
      if tail[i] % head == 0 {
        tail[i] = tail[i] / head
      }
    }
    finalMulti = append(finalMulti, head)
    multiples = tail
  }

  product := 1
  for _, multi := range finalMulti {
    product = product * multi
  }
  fmt.Println(product)
}
