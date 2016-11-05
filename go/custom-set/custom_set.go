package stringset

import (
	"strconv"
	"strings"
)

const testVersion = 3

// Set interface
type Set interface {
	Add(string)
	Delete(string)
	Has(string) bool
	IsEmpty() bool
	Len() int
	Slice() []string
	String() string
}

// -- mySet Type --------------------------------------------------------------

type mySet struct {
	val map[string]bool
}

// New ..
func New() Set {
	return &mySet{val: make(map[string]bool)}
}

// NewFromSlice ..
func NewFromSlice(source []string) Set {
	out := &mySet{val: make(map[string]bool)}
	for _, s := range source {
		out.val[s] = true
	}
	return out
}

// Add ..
func (my *mySet) Add(str string) { my.val[str] = true }

// Delete ..
func (my *mySet) Delete(str string) { delete(my.val, str) }

// Has ..
func (my *mySet) Has(str string) bool {
	_, ok := my.val[str]
	return ok
}

// IsEmpty ..
func (my *mySet) IsEmpty() bool { return len(my.val) == 0 }

// Len ..
func (my *mySet) Len() int { return len(my.val) }

// Slice ..
func (my *mySet) Slice() []string {
	var out []string
	for str := range my.val {
		out = append(out, str)
	}
	return out
}

// String ..
func (my *mySet) String() string {
	var s []string
	for str := range my.val {
		s = append(s, strconv.Quote(str))
	}
	return "{" + strings.Join(s, ", ") + "}"
}

// -- Set Functions -----------------------------------------------------------

// Equal ..
func Equal(s1, s2 Set) bool {
	if s1.Len() != s2.Len() {
		return false
	}
	for _, s := range s1.Slice() {
		if !s2.Has(s) {
			return false
		}
	}
	return true
}

// Subset ..
func Subset(s1, s2 Set) bool {
	for _, c := range s1.Slice() {
		if !s2.Has(c) {
			return false
		}
	}
	return true
}

// Disjoint ..
func Disjoint(s1, s2 Set) bool {
	for _, s := range s1.Slice() {
		if s2.Has(s) {
			return false
		}
	}
	return true
}

// Intersection ..
func Intersection(s1, s2 Set) Set {
	out := &mySet{val: make(map[string]bool)}
	for _, s := range s1.Slice() {
		if s2.Has(s) {
			out.Add(s)
		}
	}
	return out
}

// Union ..
func Union(s1, s2 Set) Set {
	out := &mySet{val: make(map[string]bool)}
	for _, s := range s1.Slice() {
		out.Add(s)
	}
	for _, s := range s2.Slice() {
		out.Add(s)
	}
	return out
}

// Difference ..
func Difference(s1, s2 Set) Set {
	out := &mySet{val: make(map[string]bool)}
	for _, s := range s1.Slice() {
		if !s2.Has(s) {
			out.Add(s)
		}
	}
	return out
}

// SymmetricDifference ..
func SymmetricDifference(s1, s2 Set) Set {
	return Union(Difference(s1, s2), Difference(s2, s1))
}
