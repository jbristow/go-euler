/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143?

https://projecteuler.net/problem=3
*/
package problem003

import "../../prime"

func Problem003(max int) int {
	primes := prime.DefaultPrimes
	maxPrime := 0
	for i := 5; i*i <= max; i += 2 {
		if max%int(i) == 0 {
			var isPrime bool
			isPrime, primes = prime.IsPrime(i, primes)
			if isPrime {
				maxPrime = i
			}
		}
	}

	return maxPrime
}
