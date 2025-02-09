package main

func climbStairs(n int) int {
	if n < 0 {
		return 1
	}
	if n < 3 {
		return n
	}

	tic, tac := 1, 2
	for i := 2; i < n; i++ {
		tic, tac = tac, tic+tac
	}

	return tac
}
