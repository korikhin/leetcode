package main

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}

	for i, k := 0, 0; k < len(t); k++ {
		if s[i] == t[k] {
			i++
		}
		if i == len(s) {
			return true
		}
	}

	return false
}
