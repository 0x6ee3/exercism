package atbash

import (
	"strings"
)

var cipher = map[rune]rune{}

func init() {
	for i := 0; i < 26; i++ {
		cipher[rune('a'+i)] = rune('z' - i)
	}
}

func mapNormalizeFunc(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + 32
	}
	if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
		return r
	}
	return -1
}

func mapCipherFunc(r rune) rune {
	if c, ok := cipher[r]; ok {
		return c
	}
	return r
}

// Atbash ..
func Atbash(in string) string {
	in = strings.Map(mapNormalizeFunc, in)
	in = strings.Map(mapCipherFunc, in)

	var out string
	for i, r := range in {
		if i%5 == 0 && i != 0 {
			out += " "
		}
		out += string(r)
	}
	return out
}
