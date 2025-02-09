package main

func rangeBitwiseAnd(left int, right int) int {
	n := right
	for n > left {
		n &= n - 1
	}

	return n
}
