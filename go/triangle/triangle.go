package triangle

import "math"

const testVersion = 3

// Kind of triangle
type Kind int

// Kinds of triangles
const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides returns the kind of triangle from given sides.
func KindFromSides(a, b, c float64) Kind {
	if math.IsNaN(a) || math.IsInf(a, 0) || a <= 0 ||
		math.IsNaN(b) || math.IsInf(b, 0) || b <= 0 ||
		math.IsNaN(c) || math.IsInf(c, 0) || c <= 0 {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a+b >= c && a+c >= b && b+c >= a {
		if a == b || a == c || b == c {
			return Iso
		}
		return Sca
	}
	return NaT
}
