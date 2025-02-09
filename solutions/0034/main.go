package main

func findStart(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if left < len(nums) && nums[left] == target {
		return left
	}
	return -1
}

func findEnd(nums []int, target, start int) int {
	left, right := start, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	if target < nums[0] || nums[len(nums)-1] < target {
		return []int{-1, -1}
	}

	start := findStart(nums, target)
	if start == -1 {
		return []int{-1, -1}
	}
	end := findEnd(nums, target, start)

	return []int{start, end}
}
