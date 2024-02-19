package main

import "math"

func Reverse(x int) int {
	r := 0
	for y := abs(x); y > 0; {
		r = 10*r + y%10
		// hacky, but c'mon...
		if r < math.MinInt32 || math.MaxInt32 < r {
			return 0
		}
		y /= 10
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
