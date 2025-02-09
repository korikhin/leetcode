package main

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Fields(s)
	if len(pattern) != len(words) {
		return false
	}

	pMap := make(map[byte]string)
	seen := make(map[string]struct{})

	for i, word := range words {
		char := pattern[i]
		pWord, exists := pMap[char]
		if exists && pWord != word {
			return false
		}
		if _, hasSeen := seen[word]; !exists && hasSeen {
			return false
		}
		pMap[char] = word
		seen[word] = struct{}{}
	}

	return true
}
