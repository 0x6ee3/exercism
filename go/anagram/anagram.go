package anagram

import (
	"strings"
)

func isAnagram(subject, candidate string) bool {
	if len(subject) != len(candidate) || subject == candidate {
		return false
	}

	for _, r := range subject {
		if strings.ContainsRune(candidate, r) {
			candidate = strings.Replace(candidate, string(r), "", 1)
		}
	}

	return len(candidate) == 0
}

func contains(arr []string, in string) bool {
	for _, a := range arr {
		if a == in {
			return true
		}
	}
	return false
}

// Detect ..
func Detect(subject string, candidates []string) []string {
	var matches []string
	subject = strings.ToLower(subject)

	for _, candidate := range candidates {
		candidate = strings.ToLower(candidate)

		if !contains(matches, candidate) && isAnagram(subject, candidate) {
			matches = append(matches, candidate)
		}
	}

	return matches
}
