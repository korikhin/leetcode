package main

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	builder := make([]byte, len(a)+1)
	index := len(a)

	for i, j, c := len(a)-1, len(b)-1, 0; i >= 0 || j >= 0 || c > 0; {
		if i >= 0 {
			c += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			c += int(b[j] - '0')
			j--
		}
		builder[index] = byte(c%2 + '0')
		index--
		c /= 2
	}

	result := string(builder[index+1:])

	if result == "" {
		return "0"
	}
	return result
}
