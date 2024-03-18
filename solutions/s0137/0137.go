package main

// func SingleNumber(nums []int) int {
// 	var ones, twos int
// 	for i := range nums {
// 		ones ^= nums[i] & ^twos
// 		twos ^= nums[i] & ^ones
// 	}
//
// 	return ones
// }

func SingleNumber(nums []int) int {
	return SingleNumberN(nums, 3 /* repetitions */)
}

func SingleNumberN(nums []int, p int) int {
	if p&1 == 0 {
		var result int
		for i := range nums {
			result ^= nums[i]
		}

		return result
	}

	var mask int
	s := make([]int, p)

	for i := range nums {
		for k := p - 1; k > 0; k-- {
			s[k] ^= s[k-1] & nums[i]
		}
		s[0] ^= nums[i]

		mask = ^(s[p-1] & s[p-2])
		for k := range s {
			s[k] &= mask
		}
	}

	return s[0]
}
