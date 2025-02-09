package main

func LengthOfLongestSubstring(s string) int {
	index := make(map[rune]int)
	var left, result int

	for right, r := range s {
		if duplicatePos, ok := index[r]; ok && duplicatePos >= left {
			left = duplicatePos + 1
		}
		result = max(result, right-left+1)
		index[r] = right
	}

	return result
}
