package main

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	if n < 2 || k < 1 {
		return
	}

	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(a []int, start, end int) {
	for start < end {
		a[start], a[end] = a[end], a[start]
		start++
		end--
	}
}
