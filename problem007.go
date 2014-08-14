package main

import (
	"./prime"
	"fmt"
)

func main() {

	primes := prime.DefaultPrimes
	for i := 5; len(primes) < 10001; i += 2 {
		_, primes = prime.IsPrime(i, primes)
	}
	fmt.Println(primes[len(primes)-1])
}
