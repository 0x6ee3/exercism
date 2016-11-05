package strain

// Ints ..
type Ints []int

// Lists ..
type Lists [][]int

// Strings ..
type Strings []string

// Keep ..
func (i Ints) Keep(f func(int) bool) Ints {
	var out Ints
	for _, v := range i {
		if f(v) {
			out = append(out, v)
		}
	}
	return out
}

// Discard ..
func (i Ints) Discard(f func(int) bool) Ints {
	var out Ints
	for _, v := range i {
		if !f(v) {
			out = append(out, v)
		}
	}
	return out
}

// Keep ..
func (l Lists) Keep(f func([]int) bool) Lists {
	var out Lists
	for _, v := range l {
		if f(v) {
			out = append(out, v)
		}
	}
	return out
}

// Keep ..
func (s Strings) Keep(f func(string) bool) Strings {
	var out Strings
	for _, v := range s {
		if f(v) {
			out = append(out, v)
		}
	}
	return out
}
