package main

func IsMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true

	// Checks if character mathces the pattern so far
	isMatchChar := func(i int, j int) bool {
		if i < 0 {
			return false
		}
		return p[j] == '.' || s[i] == p[j]
	}

	// Initialize for patterns with '*'
	for j := 1; j < n; j += 2 {
		if p[j] == '*' && dp[0][j-1] {
			dp[0][j+1] = true
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if p[j] == '*' {
				// Check for zero occurrence or more occurrences of the character before '*'
				noRepeat := dp[i+1][j-1]
				doRepeat := isMatchChar(i, j-1) && dp[i][j+1]
				dp[i+1][j+1] = noRepeat || doRepeat
			} else if isMatchChar(i, j) {
				dp[i+1][j+1] = dp[i][j]
			}
		}
	}

	return dp[m][n]
}
