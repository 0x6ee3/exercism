package lsproduct

import (
	"errors"
	"unicode"
)

const testVersion = 4

func stringIsNumber(number string) bool {
	for _, c := range number {
		if !unicode.IsNumber(c) {
			return false
		}
	}
	return true
}

// LargestSeriesProduct ..
func LargestSeriesProduct(digits string, span int) (int64, error) {
	lenDigits := len(digits)

	if span < 0 || lenDigits < span || !stringIsNumber(digits) {
		return 0, errors.New("Invalid input.")
	}
	if span == 0 {
		return 1, nil
	}

	var sum, highestSum int64
	for i := 0; i <= lenDigits-span; i++ {
		sum = 1
		for _, c := range digits[i : i+span] {
			if c == '0' {
				goto skip
			}
			sum *= int64(c - '0')
		}

		if sum > highestSum {
			highestSum = sum
		}

	skip:
	}

	return highestSum, nil
}
