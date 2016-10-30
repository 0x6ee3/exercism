package wordcount

import (
	"strings"
	"unicode"
)

const testVersion = 2

// Frequency ..
type Frequency map[string]int

// use normalizeFunc to remove any non letter or digit
func normalizeFunc(r rune) bool {
	return !unicode.IsLetter(r) && !unicode.IsDigit(r)
}

// WordCount ..
func WordCount(phrase string) Frequency {
	words := Frequency{}
	fields := strings.FieldsFunc(phrase, normalizeFunc)
	for _, v := range fields {
		words[strings.ToLower(v)]++
	}
	return words
}
