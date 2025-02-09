package main

import "strings"

func intToRoman(num int) string {
	romans := []struct {
		Value  int
		Symbol string
	}{
		{1000, "M"}, {900, "CM"},
		{500, "D"}, {400, "CD"},
		{100, "C"}, {90, "XC"},
		{50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"},
		{5, "V"}, {4, "IV"},
		{1, "I"},
	}

	var (
		c int
		m = num
		b strings.Builder
	)

	for _, r := range romans {
		if m == 0 {
			break
		}
		c = m / r.Value
		m = m % r.Value
		b.WriteString(strings.Repeat(r.Symbol, c))
	}

	return b.String()
}
