package phonenumber

import (
	"errors"
	"strings"
)

const testVersion = 1

func normalize(r rune) rune {
	if r >= '0' && r <= '9' {
		return r
	}
	return -1
}

// Number ..
func Number(input string) (string, error) {
	input = strings.Map(normalize, input)
	lenInput := len(input)

	if lenInput == 10 {
		return input, nil
	}
	if lenInput == 11 && input[:1] == "1" {
		return input[1:], nil
	}

	return "", errors.New("Invalid input")
}

// AreaCode ..
func AreaCode(input string) (string, error) {
	input, err := Number(input)
	if err != nil {
		return "", err
	}

	return input[:3], nil
}

// Format ..
func Format(input string) (string, error) {
	input, err := Number(input)
	if err != nil {
		return "", err
	}

	return "(" + input[:3] + ") " + input[3:6] + "-" + input[6:], nil
}
