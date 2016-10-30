package robotname

import (
	"fmt"
	"math/rand"
)

// Robot ..
type Robot struct {
	name string
}

func (r *Robot) generateName() {
	var name string
	name += string('A' + rand.Intn(26))
	name += string('A' + rand.Intn(26))
	name += fmt.Sprintf("%03d", rand.Intn(100))
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
