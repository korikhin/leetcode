package main

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 2; i <= n; i += 2 {
		if p[i-1] == '*' {
			dp[i] = dp[i-2]
		}
	}

	for i := 1; i <= m; i++ {
		prev := dp[0]
		dp[0] = false

		for j := 1; j <= n; j++ {
			cur := dp[j]
			if p[j-1] == '*' {
				dp[j] = dp[j-2] || (dp[j] && (s[i-1] == p[j-2] || p[j-2] == '.'))
			} else {
				dp[j] = prev && (s[i-1] == p[j-1] || p[j-1] == '.')
			}
			prev = cur
		}
	}

	return dp[n]
}
