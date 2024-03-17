package main

func HammingWeight(num uint32) int {
	w := 0
	for num > 0 {
		num &= num - 1
		w++
	}

	return w
}
