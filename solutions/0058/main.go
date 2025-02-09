package main

func lengthOfLastWord(s string) int {
	end := len(s) - 1
	for end >= 0 && s[end] == ' ' {
		end--
	}
	if end < 0 {
		return 0
	}

	start := end
	for start >= 0 && s[start] != ' ' {
		start--
	}

	return end - start
}
