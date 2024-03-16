package main

func RemoveDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	p := 2
	for i := p; i < len(nums); i++ {
		if nums[i] != nums[p-2] {
			nums[p] = nums[i]
			p++
		}
	}

	return p
}
