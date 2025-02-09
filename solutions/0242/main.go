package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := make([]int, 26)
	tMap := make([]int, 26)

	for i := range s {
		sMap[s[i]-'a']++
		tMap[t[i]-'a']++
	}

	for i := range sMap {
		if sMap[i] != tMap[i] {
			return false
		}
	}

	return true
}

// NOTE: Follow Up
// func isAnagram(s, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}

// 	chars := make(map[rune]int)

// 	for _, r := range s {
// 		chars[r]++
// 	}
// 	for _, r := range t {
// 		chars[r]--
// 	}

// 	for _, count := range chars {
// 		if count != 0 {
// 			return false
// 		}
// 	}

// 	return true
// }
