package main

import "math"

func MyAtoi(s string) int {
	i, n, a, sign := 0, len(s), 0, 1

	// ignore whitespaces
	for i < n && s[i] == ' ' {
		i++
	}
	if i == n {
		return 0
	}

	// check the sign
	if isNegative := s[i] == '-'; isNegative || s[i] == '+' {
		if isNegative {
			sign = -1
		}
		i++
	}

	for i < n && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')

		// check overflow before doing math
		if sign > 1 {
			if a > (math.MaxInt32-digit)/10 {
				return math.MaxInt32
			}
		} else {
			if -a < (math.MinInt32+digit)/10 {
				return math.MinInt32
			}
		}
		a = a*10 + digit
		i++
	}

	return a * sign
}
