package main

import "math"

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	a := 0
	for y := abs(x); y > 0; y /= 10 {
		digit := y % 10

		// check overflow before doing math
		if x > 0 && a > (math.MaxInt32-digit)/10 {
			return false
		}
		if x < 0 && -a < (math.MinInt32+digit)/10 {
			return false
		}
		a = 10*a + digit
	}

	return x == a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
