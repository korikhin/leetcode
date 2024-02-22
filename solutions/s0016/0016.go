package main

import "sort"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ThreeSumClosest(nums []int, target int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}

	sort.Ints(nums)
	sumClosest := nums[0] + nums[1] + nums[2]

	for i := 0; i < n-2; i++ {
		// skip duplicates
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]

			if sum == target {
				return sum
			}

			if abs(target-sum) < max(target-sumClosest) {
				sumClosest = sum
			}

			if sum < target {
				l++
			} else {
				r--
			}
		}
	}

	return sumClosest
}
