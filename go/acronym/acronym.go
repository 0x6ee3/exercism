package acronym

import (
	"strings"
	"unicode"
)

const testVersion = 1

// wordFieldFunc treads everything that isn't a letter as a field seperator.
func wordFieldFunc(r rune) bool {
	return !unicode.IsLetter(r)
}

// isUpperCaseWord checks if the word consists only of upper case characters.
func isUpperCaseWord(word string) bool {
	for _, r := range word {
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

// abbreviate generates a short acronym.
func abbreviate(name string) string {
	var abb string

	for _, word := range strings.FieldsFunc(name, wordFieldFunc) {
		// Only use the first letter on all upper case words
		if isUpperCaseWord(word) {
			abb += word[:1]
		} else {
			// Use the first character or any upper case character
			for i, w := range word {
				if i == 0 || unicode.IsUpper(w) {
					abb += string(w)
				}
			}
		}
	}

	return strings.ToUpper(abb)
}
