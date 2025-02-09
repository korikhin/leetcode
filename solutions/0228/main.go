package main

import (
	"fmt"
)

func summaryRanges(nums []uint64) (r []string) {
	for i, k := 0, 0; k < len(nums); k++ {
		i = k
		for k < len(nums)-1 && nums[k] == nums[k+1]-1 {
			k++
		}

		if i == k {
			r = append(r, fmt.Sprint(nums[i]))
		} else {
			r = append(r, fmt.Sprintf("%d->%d", nums[i], nums[k]))
		}
	}

	return r
}
