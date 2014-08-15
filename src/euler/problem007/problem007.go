package euler

import "../prime"

func Problem007() int {
	primes := prime.DefaultPrimes
	for i := 5; len(primes) < 10001; i += 2 {
		_, primes = prime.IsPrime(i, primes)
	}
	return primes[len(primes)-1]
}
