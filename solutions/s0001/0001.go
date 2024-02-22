package main

func TwoSum(nums []int, target int) []int {
	n := len(nums)
	if n < 2 {
		return nil
	}

	m := make(map[int]int, n)

	for i, num := range nums {
		if j, ok := m[target-num]; ok {
			return []int{j, i}
		}
		m[n] = i
	}

	return nil
}
