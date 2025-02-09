package main

type MinStack struct {
	s []int
	m int
}

func Constructor() MinStack {
	return MinStack{s: make([]int, 0)}
}

func (ms *MinStack) Push(x int) {
	if len(ms.s) == 0 {
		ms.s = append(ms.s, x)
		ms.m = x
	} else if x < ms.m {
		ms.s = append(ms.s, 2*x-ms.m)
		ms.m = x
	} else {
		ms.s = append(ms.s, x)
	}
}

func (ms *MinStack) Pop() {
	n := len(ms.s)
	if n == 0 {
		return
	}
	if top := ms.s[n-1]; top < ms.m {
		ms.m = 2*ms.m - top
	}
	ms.s = ms.s[:n-1]
}

func (ms *MinStack) Top() int {
	n := len(ms.s)
	if n == 0 {
		return 0
	}
	if top := ms.s[n-1]; top < ms.m {
		return ms.m
	} else {
		return top
	}
}

func (ms *MinStack) Min() int {
	return ms.m
}
