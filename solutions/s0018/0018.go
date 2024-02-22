package main

import "sort"

func FourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return nSum(nums, target, 4 /* depth */, 0 /* start */)
}

func nSum(nums []int, target, n, start int) [][]int {
	var result [][]int

	// Not enough elements
	if n < 2 || n > len(nums)-start {
		return nil
	}

	// Target is out of rich
	if target < nums[start]*n || target > nums[len(nums)-1]*n {
		return nil
	}

	// Base case
	if n == 2 {
		left, right := start, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum < target {
				left++
				continue
			}
			if sum > target {
				right--
				continue
			}
			result = append(result, []int{nums[left], nums[right]})
			left++
			right--
			for left < right && nums[left] == nums[left-1] {
				left++
			}
			for left < right && nums[right] == nums[right+1] {
				right--
			}
		}

		return result
	}

	for i := start; i <= len(nums)-n; i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		subResults := nSum(nums, target-nums[i], n-1, i+1)
		for _, sub := range subResults {
			result = append(result, append([]int{nums[i]}, sub...))
		}
	}

	return result
}
