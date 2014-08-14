package main

import (
	"fmt"
)

func main() {
  listA := []int {1,1}
  for i:=2; i<=40; i++ {
    listB := []int {1}
    for j := 1; j < len(listA); j++ {
      listB = append(listB, listA[j]+listA[j-1])
    }
    listA = append(listB, 1)
  }
  fmt.Println(listA[len(listA)/2])
}
