package main

import (
	"fmt"
)

func main() {
	maxCount := 0
	maxNum := 0
	for i := 3; i < 1000000; i++ {
		count := 0
		num := i
		for num > 1 {
			count += 1
			num = collatz(num)
		}
		if count > maxCount {
			maxCount = count
			maxNum = i
		}
	}
	fmt.Println(maxNum)
}

func collatz(num int) (next int) {
	if num%2 == 0 {
		next = num / 2
	} else {
		next = 3*num + 1
	}
	return
}
