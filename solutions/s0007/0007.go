package main

import "math"

func Reverse(x int) int {
	r := 0
	for y := abs(x); y > 0; y /= 10 {
		digit := y % 10

		// check overflow before doing math
		if x > 0 && r > (math.MaxInt32-digit)/10 {
			return 0
		}
		if x < 0 && -r < (math.MinInt32+digit)/10 {
			return 0
		}
		r = 10*r + digit
	}

	if x < 0 {
		return -r
	}
	return r
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
