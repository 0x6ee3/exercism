package bob

import (
	"strings"
	"unicode"
)

const testVersion = 2

func isSilence(line string) bool {
	for _, v := range line {
		if !unicode.IsSpace(v) {
			return false
		}
	}
	return true
}

func isQuestion(line string) bool {
	line = strings.TrimSpace(line)
	return line[len(line)-1:] == "?"
}

func isShouted(line string) bool {
	var atLeastOneLetter bool
	for _, v := range line {
		if unicode.IsLetter(v) {
			atLeastOneLetter = true
			if !unicode.IsUpper(v) {
				return false
			}
		}
	}
	return atLeastOneLetter
}

// Hey askes Bob a question.
func Hey(line string) string {
	if isSilence(line) {
		return "Fine. Be that way!"
	}
	if isShouted(line) {
		return "Whoa, chill out!"
	}
	if isQuestion(line) {
		return "Sure."
	}

	return "Whatever."
}
