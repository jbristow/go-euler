package main

import (
  "fmt"
  "math"
  "./prime"
)


const (
  MAXNUM int64 = 600851475143
)

func main() {
  primes := prime.DefaultPrimes
  maxPrime := 0;
  for i :=5; float64(i)<math.Sqrt(float64(MAXNUM)); i+=2 {
    if MAXNUM % int64(i) == 0 {
      var isPrime bool
      isPrime, primes = prime.IsPrime(i, primes)  
      if isPrime {
        maxPrime = i
      }
    }
  }
  fmt.Println(maxPrime)
}
