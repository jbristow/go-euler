package euler

func Problem009() int {
	for a := 1; a < 1000; a++ {
		for b := a; b < 1000-a; b++ {
			c := 1000 - a - b
			if c*c == a*a+b*b {
				return a * b * c
			}
		}
	}
	return 0
}
