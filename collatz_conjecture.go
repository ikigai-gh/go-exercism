package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return -1, errors.New("Argument can not be <= 0")
	}
	steps := 0

	for {
		if n == 1 {
			break
		}
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		steps++
	}
	return steps, nil
}
