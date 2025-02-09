package main

func trailingZeroes(n int) int {
	if n < 0 {
		return -1
	}

	count := 0
	for n >= 5 {
		n /= 5
		count += n
	}
	return count
}
