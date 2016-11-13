package dna

import "errors"

const testVersion = 2

// Histogram ..
type Histogram map[byte]int

// DNA ..
type DNA string

func isValid(r byte) bool {
	return r == 'A' || r == 'C' || r == 'T' || r == 'G'
}

// Count ..
func (d DNA) Count(r byte) (int, error) {
	if !isValid(r) {
		return 0, errors.New("Invalid input")
	}

	var c int
	for _, v := range d {
		if r == byte(v) {
			c++
		}
	}

	return c, nil
}

// Counts ..
func (d DNA) Counts() (Histogram, error) {
	values := Histogram{'A': 0, 'C': 0, 'T': 0, 'G': 0}

	for _, r := range d {
		if isValid(byte(r)) {
			values[byte(r)]++
		} else {
			return nil, errors.New("Invalid input")
		}
	}

	return values, nil
}
