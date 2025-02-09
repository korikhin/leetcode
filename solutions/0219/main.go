package main

func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) < 2 || k <= 0 {
		return false
	}

	m := make(map[int]int)
	for i := range nums {
		if p, exists := m[nums[i]]; exists && i-p <= k {
			return true
		}
		m[nums[i]] = i
	}

	return false
}
