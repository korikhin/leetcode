package main

func longestConsecutive(nums []int) (r int) {
	m := make(map[int]struct{}, len(nums))
	for _, val := range nums {
		m[val] = struct{}{}
	}

	for n := range m {
		if _, exists := m[n-1]; exists {
			continue
		}

		i := 1
		for {
			if _, exists := m[n+i]; !exists {
				break
			}
			i++
		}
		if i > r {
			r = i
		}
	}

	return r
}
