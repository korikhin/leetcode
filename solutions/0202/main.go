package main

func isHappy(n int) bool {
	m := make(map[int]struct{})
	for {
		n = reduce(n)
		if n == 1 {
			return true
		}
		if _, ok := m[n]; ok {
			return false
		} else {
			m[n] = struct{}{}
		}
	}
}

func reduce(a int) (b int) {
	for a > 0 {
		d := a % 10
		b += d * d
		a /= 10
	}
	return b
}
