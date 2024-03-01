package main

import (
	"strconv"
	"strings"
)

func CountAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	s := "1"
	for i := 2; i <= n; i++ {
		var b strings.Builder
		count := 1
		for j := 1; j < len(s); j++ {
			if s[j] == s[j-1] {
				count++
				continue
			}
			b.WriteString(strconv.Itoa(count))
			b.WriteByte(s[j-1])
			count = 1
		}
		b.WriteString(strconv.Itoa(count))
		b.WriteByte(s[len(s)-1])
		s = b.String()
	}

	return s
}
