package main

func IsValid(s string) bool {
	if len(s)&1 == 1 {
		return false
	}

	pairs := map[rune]rune{')': '(', '}': '{', ']': '['}
	var stack []rune

	for _, r := range s {
		if opening, ok := pairs[r]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != opening {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, r)
		}
	}

	return len(stack) == 0
}
