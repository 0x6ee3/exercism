package say

import "strings"

var quantifier = []string{"", " thousand", " million", " billion", " trillion", " quadrillion", " quintillion"}
var smallNums = map[int]string{
	0: "", 1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six", 7: "seven", 8: "eight", 9: "nine",
	10: "ten", 11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen", 15: "fifteen", 16: "sixteen", 17: "seventeen", 18: "eighteen", 19: "nineteen",
}
var doubleDigit = map[int]string{1: "one", 2: "twenty", 3: "thirty", 4: "forty", 5: "fifty", 6: "sixty", 7: "seventy", 8: "eighty", 9: "ninety"}

func numToString(num int) string {
	var words []string

	if num > 99 {
		words = append(words, smallNums[num/100]+" hundred")
		num %= 100
	}

	if num > 0 {
		if num < 20 {
			words = append(words, smallNums[num])
		} else {
			word := doubleDigit[num/10]
			if num%10 != 0 {
				word += "-" + smallNums[num%10]
			}
			words = append(words, word)
		}
	}

	return strings.Join(words, " ")
}

// Say ..
func Say(num uint64) string {
	if num == 0 {
		return "zero"
	}

	var numWords []string
	var i int

	for num > 0 {
		tripple := num % 1000
		trippleWord := numToString(int(tripple))

		if len(trippleWord) > 0 {
			trippleWord += quantifier[i]
			numWords = append(numWords, trippleWord)
		}

		i++
		num -= tripple
		if num == 0 {
			break
		}
		num /= 1000
	}

	var reversed []string
	for i := len(numWords) - 1; i >= 0; i-- {
		if len(numWords[i]) > 0 {
			reversed = append(reversed, numWords[i])
		}
	}

	return strings.Join(reversed, " ")
}
