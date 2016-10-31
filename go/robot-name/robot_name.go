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

func (r *Robot) generateName() {
	var name string
	name += string('A' + rng.Intn(26))
	name += string('A' + rng.Intn(26))
	name += fmt.Sprintf("%03d", rng.Intn(999))
	r.name = name
}

// Name ..
func (r *Robot) Name() string {
	if r.name == "" {
		r.generateName()
	}

	return r.name
}

// Reset ..
func (r *Robot) Reset() {
	r.name = ""
}

func init() {

}
