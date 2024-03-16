package main

func IsMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := range p {
		if p[i] == '*' {
			dp[i+1] = dp[i]
		}
	}

	for i := 1; i <= m; i++ {
		prev := dp[0]
		dp[0] = false

		for j := 1; j <= n; j++ {
			cur := dp[j]
			if p[j-1] == '*' {
				dp[j] = dp[j-1] || dp[j] || prev
			} else {
				dp[j] = prev && (s[i-1] == p[j-1] || p[j-1] == '?')
			}
			prev = cur
		}
	}

	return dp[n]
}
