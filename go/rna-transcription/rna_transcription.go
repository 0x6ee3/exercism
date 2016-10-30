package strand

import "strings"

const testVersion = 3

func dnaToRna(r rune) rune {
	if r == 'C' {
		return 'G'
	}
	if r == 'G' {
		return 'C'
	}
	if r == 'T' {
		return 'A'
	}
	if r == 'A' {
		return 'U'
	}
	return r
}

// ToRNA ..
func ToRNA(dna string) string {
	return strings.Map(dnaToRna, dna)
}
