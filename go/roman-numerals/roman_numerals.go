package romannumerals

import (
	"errors"
	"strings"
)

const testVersion = 3

var romanTable = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

func largestNumeral(num int) (int, string) {
	var largestNum int
	var largestRoman string
	for i, v := range romanTable {
		if i > largestNum && i <= num {
			largestRoman = v
			largestNum = i
		}
	}
	return largestNum, largestRoman
}

// ToRomanNumeral ..
func ToRomanNumeral(num int) (string, error) {
	if num < 1 || num >= 4000 {
		return "", errors.New("Input out of range!")
	}

	var nums []string
	for num > 0 {
		value, roman := largestNumeral(num)
		nums = append(nums, roman)
		num -= value
	}
	return strings.Join(nums, ""), nil
}
