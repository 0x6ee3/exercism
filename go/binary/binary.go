package binary

import "errors"

const testVersion = 2

func isBinaryString(input string) bool {
	for _, c := range input {
		if c != '0' && c != '1' {
			return false
		}
	}
	return true
}

// ParseBinary ..
func ParseBinary(input string) (int, error) {
	if !isBinaryString(input) {
		return 0, errors.New("Invalid input")
	}

	var sum int
	for i, j := len(input)-1, 0; i >= 0; i, j = i-1, j+1 {
		if input[i] == '1' {
			sum += 1 << uint(j)
		}
	}
	return sum, nil
}
