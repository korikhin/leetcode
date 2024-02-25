package main

func SearchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	l, r := 0, len(nums)
	for l < r {
		m := (l + r) / 2
		if target == nums[m] {
			return m
		}
		if target > nums[m] {
			l = m + 1
		} else {
			r = m
		}
	}

	return l
}
