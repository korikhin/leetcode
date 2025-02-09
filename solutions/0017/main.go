package main

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	keys := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var (
		a   []string
		dfs func(string, string)
	)

	dfs = func(c, next string) {
		if len(next) == 0 {
			a = append(a, c)
			return
		}

		for _, r := range keys[next[0]] {
			c := c + string(r)
			dfs(c, next[1:])
		}
	}

	dfs("", digits)
	return a
}
