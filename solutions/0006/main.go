package main

func convert(s string, numRows int) string {
	n := len(s)
	if numRows == 1 || numRows >= n {
		return s
	}

	z := make([]byte, n)
	step, count := (numRows-1)<<1, 0

	for i := 0; i < numRows; i++ {
		for j := 0; j+i < n; j += step {
			z[count] = s[i+j]
			count++

			if i > 0 && i != numRows-1 && j+step-i < n {
				z[count] = s[j+step-i]
				count++
			}
		}
	}

	return string(z)
}
