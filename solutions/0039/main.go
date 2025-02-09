package main

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var (
		result [][]int
		crunch func(int, int, []int)
	)

	crunch = func(start, target int, combination []int) {
		if target == 0 {
			result = append(result, append([]int(nil), combination...))
			return
		}

		for i := start; i < len(candidates); i++ {
			if candidates[i] > target {
				break
			}
			crunch(i, target-candidates[i], append(combination, candidates[i]))
		}
	}

	crunch(0, target, nil)
	return result
}
