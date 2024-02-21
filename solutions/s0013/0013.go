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
		if romans[s[i]] < romans[s[i+1]] {
			a -= romans[s[i]]
		} else {
			a += romans[s[i]]
		}
	}

	return a
}
