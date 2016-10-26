package slice

// All ...
func All(n int, s string) []string {
	if n > len(s) || n <= 0 {
		return nil
	}
	if n == len(s) {
		return []string{s}
	}

	matches := len(s) - n + 1
	all := make([]string, matches)
	for i := 0; i < matches; i++ {
		all[i] = s[i : i+n]
	}

	return all
}

// UnsafeFirst ...
func UnsafeFirst(n int, s string) string {
	return s[:n]
}

// First ...
func First(n int, s string) (first string, ok bool) {
	if n > len(s) || n <= 0 {
		return "", false
	}
	return UnsafeFirst(n, s), true
}
