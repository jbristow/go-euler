package main

import (
	"fmt"
)

func main() {

	sum := 0
	divisorCount := 0
	for i := 1; divisorCount <= 500; i++ {
		sum += i
		divisors := findDivisors(sum)
		divisorCount = len(divisors)
	}
	fmt.Println(sum)
}

func findDivisors(num int) []int {
	divisors := []int{}
	for i := 1; i*i < num; i++ {
		if num%i == 0 {
			divisors = append(divisors, i, num/i)
		}

	}
	return divisors
}
