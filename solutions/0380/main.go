package main

import "math/rand"


type RandomizedSet struct {
	m map[int]int
	a []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		m: make(map[int]int),
		a: make([]int, 0),
	}
}

func (s *RandomizedSet) Insert(val int) bool {
	if _, ok := s.m[val]; !ok {
		s.a = append(s.a, val)
		s.m[val] = len(s.a) - 1
		return true
	}

	return false
}

func (s *RandomizedSet) Remove(val int) bool {
	if k, ok := s.m[val]; ok {
		// Replace the removed element with the last one
		lastVal := s.a[len(s.a)-1]
		s.a[k] = lastVal
		s.m[lastVal] = k

		// Finally remove the element and shrink the slice
		s.a = s.a[:len(s.a)-1]
		delete(s.m, val)

		return true
	}

	return false
}

func (s *RandomizedSet) GetRandom() int {
	if n := len(s.a); n == 0 {
		return 0
	} else {
		return s.a[rand.Intn(n)]
	}
}
