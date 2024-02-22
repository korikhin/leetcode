package main

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	p := strs[0]
	for i := 0; i < len(p); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != p[i] {
				return p[:i]
			}
		}
	}
	return p
}
