package main

func IsMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([]bool, n+1)
	dp[0] = true

	for j := 2; j <= n; j += 2 {
		if p[j-1] == '*' {
			dp[j] = dp[j-2]
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
