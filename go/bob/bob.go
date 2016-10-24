package bob

import (
	"strings"
	"unicode"
)

const testVersion = 2

// Manually iterating over the line is ~3x faster (~2700ns/op to ~9200ns/op)
// then a compairsion like "line == line.ToUpper && hasAnyLetter(line)"
func isShouting(line string) bool {
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
	lineTrimmed := strings.TrimSpace(line)

	if len(lineTrimmed) == 0 {
		return "Fine. Be that way!"
	}
	if isShouting(lineTrimmed) {
		return "Whoa, chill out!"
	}
	if lineTrimmed[len(lineTrimmed)-1:] == "?" {
		return "Sure."
	}

	return "Whatever."
}
