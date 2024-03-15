package main

import "strings"

func Multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	result := make([]byte, len(num1)+len(num2))

	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			p := i + j + 1
			sum := result[p] + (num1[i]-'0')*(num2[j]-'0')

			result[p] = sum % 10
			result[p-1] += sum / 10
		}
	}

	var b strings.Builder
	b.Grow(len(result))

	for i, d := range result {
		if i == 0 && d == 0 {
			continue
		}
		b.WriteByte(d + '0')
	}

	return b.String()
}
