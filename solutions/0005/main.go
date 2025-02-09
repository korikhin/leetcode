package main

import "strings"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func prepare(s string) string {
	var b strings.Builder
	b.Grow(2*len(s) + 3)

	b.WriteRune('^')
	for _, r := range s {
		b.WriteRune('#')
		b.WriteRune(r)
	}
	b.WriteRune('#')
	b.WriteRune('$')
	return b.String()
}

// Manachar's Algorithm
func longestPalindrome(s string) string {
	m := prepare(s)
	P := make([]int, len(m))
	r, c := 0, 0

	for i := 1; i < len(m)-1; i++ {
		if r > i {
			P[i] = min(r-i, P[2*c-i])
		}

		for m[i+1+P[i]] == m[i-1-P[i]] {
			P[i]++
		}

		if i+P[i] > r {
			c, r = i, i+P[i]
		}
	}

	var pCenter, pMax int
	for i := 1; i < len(m)-1; i++ {
		if P[i] > pMax {
			pCenter, pMax = i, P[i]
		}
	}

	start := (pCenter - pMax) / 2
	return s[start : start+pMax]
}
