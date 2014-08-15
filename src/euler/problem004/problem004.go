package euler

import "strconv"

func isPal(s string) bool {
	mid := len(s) / 2
	last := len(s) - 1
	for i := 0; i < mid; i++ {
		if s[i] != s[last-i] {
			return false
		}
	}
	return true
}

func Problem004() int {
	maxProduct := 0
	for i := 100; i <= 999; i++ {
		for j := i; j <= 999; j++ {
			product := i * j
			if product > maxProduct && isPal(strconv.Itoa(product)) {
				maxProduct = product
			}
		}
	}
	return maxProduct
}
