package main

func MyPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}

	isNegative := n < 0
	if isNegative {
		n = -n
	}

	a := 1.0
	for m := x; n > 0; n >>= 1 {
		if n&1 != 0 {
			a *= m
		}
		m *= m
	}

	if isNegative {
		return 1 / a
	}
	return a
}
