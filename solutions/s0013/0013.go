package main

func RomanToInt(s string) int {
	var romans = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	a := romans[s[len(s)-1]]

	for i := len(s) - 2; i >= 0; i-- {
		if c := romans[s[i]]; c < romans[s[i+1]] {
			a -= c
		} else {
			a += c
		}
	}

	return a
}
