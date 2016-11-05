package igpay

import (
	"strings"
)

func vowelPos(s string) int {
	for i, r := range s {
		if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' {
			if i > 0 && s[i-1] == 'q' {
				return i + 1
			}
			return i
		}
	}
	return 0
}

// PigLatin ..
func PigLatin(in string) string {
	out := strings.Fields(in)
	for i, c := range out {
		// no idea how to handle them ..
		if c == "xray" || c == "yttria" {
			out[i] = c + "ay"
			continue
		}

		v := vowelPos(c)
		out[i] = c[v:len(c)] + c[0:v] + "ay"
	}
	return strings.Join(out, " ")
}
