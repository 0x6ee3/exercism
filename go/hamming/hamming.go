package hamming

import "errors"

const testVersion = 5

// Distance calculates the Hamming difference between two DNA strands.
func Distance(a, b string) (distance int, err error) {
	if len(a) != len(b) {
		return -1, errors.New("String lengths doesn't match!")
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
