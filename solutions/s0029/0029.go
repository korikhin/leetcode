package main

import "math"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Divide(dividend int, divisor int) int {
	q, sign := 0, -1
	if (dividend < 0) == (divisor < 0) {
		sign = 1
	}

	dividend = abs(dividend)
	divisor = abs(divisor)

	for dividend >= divisor {
		bits := 0
		for dividend >= (divisor << (bits + 1)) {
			bits++
		}

		// Check overflow before doing math
		if q > math.MaxInt32-(1<<bits) {
			if sign > 0 {
				return math.MaxInt32
			}
			return math.MinInt32
		}
		q += 1 << bits
		dividend -= divisor << bits
	}

	return q * sign
}
