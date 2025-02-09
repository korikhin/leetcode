package main

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}

	haystackRunes := []rune(haystack)
	needleRunes := []rune(needle)

	for i := 0; i <= len(haystackRunes)-len(needleRunes); i++ {
		matchFound := true
		for j := 0; j < len(needleRunes); j++ {
			if haystackRunes[i+j] != needleRunes[j] {
				matchFound = false
				break
			}
		}

		if matchFound {
			return i
		}
	}

	return -1
}
