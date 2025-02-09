package main

import "math/bits"

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	b := bits.Len(uint(x))
	i, j := 1, 1<<b/2
	for i <= j {
		m := i + (j-i)/2
		if m*m == x {
			return m
		}
		if m*m < x {
			i = m + 1
		} else {
			j = m - 1
		}
	}

	return j
}
