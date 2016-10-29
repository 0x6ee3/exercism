package luhn

import "unicode"

func normalize(str string) string {
	var norm, reverse string
	for _, c := range str {
		if unicode.IsNumber(c) {
			norm += string(c)
		}
	}
	for i := len(norm) - 1; i >= 0; i-- {
		reverse += string(norm[i])
	}

	return reverse
}

func calcCheck(in string) string {
	var sum int

	for i := 0; i < len(in); i++ {
		x := int(in[i]) - '0'

		if i != 0 && (i+1)%2 == 0 {
			sum += x
			continue
		}

		x *= 2
		if x > 9 {
			x -= 9
		}
		sum += x
	}

	sum = (sum % 10)

	if sum == 0 {
		return "0"
	}
	return string('0' + (10 - sum))
}

// Valid ..
func Valid(code string) bool {
	norm := normalize(code)

	if len(norm) == 0 {
		return false
	}

	checkString := norm[1:]
	digit := norm[:1]
	calc := calcCheck(checkString)

	return digit == calc
}

// AddCheck ..
func AddCheck(code string) string {
	norm := normalize(code)
	return code + calcCheck(norm)
}
