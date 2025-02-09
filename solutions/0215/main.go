package main

func findKthLargest(nums []int, k int) int {
	k = len(nums) - k

	var quickSelect func(l, r int) int
	quickSelect = func(l, r int) int {
		s, p := nums[r], l
		for i := l; i < r; i++ {
			if nums[i] <= s {
				nums[i], nums[p] = nums[p], nums[i]
				p++
			}
		}
		nums[r], nums[p] = nums[p], nums[r]

		if p > k {
			return quickSelect(l, p-1)
		} else if p < k {
			return quickSelect(p+1, r)
		} else {
			return nums[p]
		}
	}

	return quickSelect(0, len(nums)-1)
}
