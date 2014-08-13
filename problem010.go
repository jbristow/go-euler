package main

import (
  "fmt"
  "./prime"
)

func main() {
  primes := prime.DefaultPrimes
  sum := 2+3
  for i:=5; i<2000000; i+=2 {
    var isPrime bool
    isPrime, primes = prime.IsPrime(i, primes)
    if isPrime {
      sum += i
    }
  }
  fmt.Println(sum)
}
