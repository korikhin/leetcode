package main

import "strings"

func IntToRoman(num int) string {
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
		r strings.Builder
	)

	for _, pair := range romans {
		if m == 0 {
			break
		}
		c, m = m/pair.Value, m%pair.Value
		r.WriteString(strings.Repeat(pair.Symbol, c))
	}

	return r.String()
}
