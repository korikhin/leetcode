package main

func majorityElement(nums []int) int {
	var p int
	for c, i := 0, 0; i < len(nums); i++ {
		if c == 0 {
			p = nums[i]
		}
		if p == nums[i] {
			c++
		} else {
			c--
		}
	}

	return p
}
