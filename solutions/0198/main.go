package main

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	b, a := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		c := max(a, b+nums[i])
		b = a
		a = c
	}

	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
