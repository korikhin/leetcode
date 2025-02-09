package main

import "math"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	a := 0
	for y := abs(x); y > 0; y /= 10 {
		digit := y % 10

		// Check overflow before doing math
		if a > (math.MaxInt32-digit)/10 {
			return false
		}
		a = 10*a + digit
	}

	return x == a
}
