package main

func IsIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapS := make([]byte, 128 /* ASCII */)
	mapT := make([]byte, 128 /* ASCII */)

	for i := range s {
		sChar := s[i]
		tChar := t[i]

		if mapS[sChar] != mapT[tChar] {
			return false
		}

		// Add one to avoid confusion with \u0000 characters
		mapS[sChar] = byte(i + 1)
		mapT[tChar] = byte(i + 1)
	}

	return true
}
