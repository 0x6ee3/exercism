package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

// seeded random source (time-based)
var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// Robot ..
type Robot struct {
	name string
}

// Name ..
func (r *Robot) Name() string {
	if r.name == "" {
		r.name = string('A' + rng.Intn(25))
		r.name += string('A' + rng.Intn(25))
		r.name += fmt.Sprintf("%03d", rng.Intn(999))
	}
	return r.name
}

// Reset ..
func (r *Robot) Reset() {
	r.name = ""
}
