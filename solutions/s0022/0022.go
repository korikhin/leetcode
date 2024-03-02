package main

func GenerateParenthesis(n int) []string {
	if n < 1 {
		return nil
	}

	var (
		a []string

		// Recursively constructs parentheses combinations by tracking
		// the number of open (`opn`) and close (`cls`) brackets left to use.
		dfs func(string, int, int)
	)

	dfs = func(p string, opn, cls int) {
		if opn == 0 && cls == 0 {
			a = append(a, p)
		}
		if opn > 0 {
			p := p + "("
			dfs(p, opn-1, cls)
		}
		if opn < cls {
			p := p + ")"
			dfs(p, opn, cls-1)
		}
	}

	dfs("", n, n)
	return a
}
