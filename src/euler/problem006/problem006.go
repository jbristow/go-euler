package euler

func Problem006() int {
	sumOfSquares := 0
	sums := 0

	for i := 1; i <= 100; i++ {
		sumOfSquares += i * i
		sums += i
	}

	return sums*sums - sumOfSquares
}
