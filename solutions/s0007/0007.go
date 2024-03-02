package main

import "math"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Reverse(x int) int {
	r := 0
	for y := abs(x); y > 0; y /= 10 {
		digit := y % 10

		// Check overflow before doing math
		if r > (math.MaxInt32-digit)/10 {
			return 0
		}
		r = 10*r + digit
	}

	if x < 0 {
		return -r
	}
	return r
}
