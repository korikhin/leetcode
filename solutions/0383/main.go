package main

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]int, 26 /* English alphabet */)
	for _, r := range magazine {
		m[r]++
	}

	for _, r := range ransomNote {
		if m[r] > 0 {
			m[r]--
		} else {
			return false
		}
	}

	return true
}
