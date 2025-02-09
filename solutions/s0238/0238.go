package main

func productExceptSelf(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}

	a := make([]int, n)
	for p, i := 1, 0; i < n; i++ {
		a[i] = p
		p *= nums[i]
	}
	for p, i := 1, n-1; i >= 0; i-- {
		a[i] *= p
		p *= nums[i]
	}

	return a
}
