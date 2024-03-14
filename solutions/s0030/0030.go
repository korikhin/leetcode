package main

func FindSubstring(s string, words []string) []int {
	if len(words) == 0 || s == "" {
		return nil
	}

	k, n := len(words), len(words[0])
	if len(s) < n*k {
		return nil
	}

	wordCounts := make(map[string]int, k)
	for _, w := range words {
		wordCounts[w]++
	}

	var result []int
	queue := make([]string, 0, k)
	seenTimes := make(map[string]int, len(words))

	reset := func() {
		for k := range seenTimes {
			delete(seenTimes, k)
		}
	}

	for i := 0; i < n; i++ {
		reset()
		queue = queue[:0]
		head, qlen := 0, 0

		for j := i; j+n <= len(s); j += n {
			w := s[j : j+n]
			if count, exists := wordCounts[w]; exists {
				for seenTimes[w] == count {
					seenTimes[queue[head]]--
					head++
					qlen--
				}
				seenTimes[w]++
				queue = append(queue, w)
				qlen++
				if qlen == k {
					result = append(result, j-(k-1)*n)
				}
			} else {
				reset()
				queue = queue[:0]
				head, qlen = 0, 0
			}
		}
	}

	return result
}
