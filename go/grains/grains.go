package grains

import "errors"

const testVersion = 1

// Square ...
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("Invalid value for n")
	}

	return 1 << uint(n-1), nil
}

// Total ...
func Total() uint64 {
	return 1<<64 - 1
}
