package collatzconjecture

import "errors"

func CollatzConjecture(input int) (int, error) {
	steps := 0
	if input <= 0 {
		return -1, errors.New("Je to v hajzlu!")
	}
	for input > 1 {
		if input % 2 == 0 {
			input /= 2
		} else {
			input = 3 * input + 1
		}
		steps++
	}
	return steps, nil
}