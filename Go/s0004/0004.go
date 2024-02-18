package main

import "math"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n, m := len(nums1), len(nums2)
	if n < m {
		return FindMedianSortedArrays(nums2, nums1)
	}

	left, right := 0, m*2
	for left <= right {
		mid2 := (left + right) / 2
		mid1 := n + m - mid2

		left1, left2 := math.Inf(-1), math.Inf(-1)
		if mid1 != 0 {
			left1 = float64(nums1[(mid1-1)/2])
		}
		if mid2 != 0 {
			left2 = float64(nums2[(mid2-1)/2])
		}

		right1, right2 := math.Inf(1), math.Inf(1)
		if mid1 != n*2 {
			right1 = float64(nums1[mid1/2])
		}
		if mid2 != m*2 {
			right2 = float64(nums2[mid2/2])
		}

		if left1 > right2 {
			left = mid2 + 1
		} else if left2 > right1 {
			right = mid2 - 1
		} else {
			return (max(left1, left2) + min(right1, right2)) / 2
		}
	}
	return -1
}
