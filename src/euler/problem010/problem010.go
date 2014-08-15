package euler

import "../prime"

func problem010() int {
	primes := prime.DefaultPrimes
	sum := 2 + 3
	for i := 5; i < 2000000; i += 2 {
		var isPrime bool
		isPrime, primes = prime.IsPrime(i, primes)
		if isPrime {
			sum += i
		}
	}
	return sum
}
