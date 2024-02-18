package main

import "strings"

// Manachar's Algorithm
func LongestPalindrome(s string) string {
	m := newManacharsString(s)
	mLen := len(m)
	mSlice := make([]int, mLen)
	r, c := 0, 0

	for i := 1; i < mLen-1; i++ {
		if r > i {
			mSlice[i] = min(r-i, mSlice[2*c-i])
		}

		for m[i+1+mSlice[i]] == m[i-1-mSlice[i]] {
			mSlice[i]++
		}

		if i+mSlice[i] > r {
			c, r = i, i+mSlice[i]
		}
	}

	var pCenter, pMax int
	for i := 1; i < mLen-1; i++ {
		if mSlice[i] > pMax {
			pCenter, pMax = i, mSlice[i]
		}
	}

	return s[(pCenter-pMax)>>1 : (pCenter+pMax)>>1]
}

func newManacharsString(s string) string {
	var builder strings.Builder
	builder.Grow(len(s)*2 + 3)

	builder.WriteString("^")
	for _, r := range s {
		builder.WriteString("#")
		builder.WriteString(string(r))
	}
	builder.WriteString("#$")
	return builder.String()
}
