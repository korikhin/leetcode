package main

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	i := 0
	for k := 1; k < len(nums); k++ {
		if nums[k] != nums[i] {
			i++
			nums[i] = nums[k]
		}
	}

	return i + 1
}
