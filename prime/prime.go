package prime

var DefaultPrimes = []int{2,3}

func IsPrime(candidate int, primes []int) (bool, []int) {
  for _,seenPrime := range primes {
    if candidate % seenPrime == 0 {
      return false, primes
    }
  }
  return true, append(primes, candidate)
}
