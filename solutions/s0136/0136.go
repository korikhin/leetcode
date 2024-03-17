package main

func SingleNumber(nums []int) int {
	var result int
	for i := range nums {
		result ^= nums[i]
	}

	return result
}
