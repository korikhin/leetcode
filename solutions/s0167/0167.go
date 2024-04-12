package main

func TwoSum(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return nil
	}

	for i, k := 0, len(numbers)-1; i < k; {
		s := numbers[i] + numbers[k]
		if s == target {
			return []int{i + 1, k + 1}
		}
		if s < target {
			i++
		} else {
			k--
		}
	}

	return nil
}
