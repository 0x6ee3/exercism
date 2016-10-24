package clock

import "fmt"

const (
	testVersion   = 4
	minutesPerDay = 24 * 60
)

// Clock type.
type Clock int

// NewClock creates a Clock from given minutes
func NewClock(minutes int) Clock {
	c := minutes % minutesPerDay

	// "Invert" negative minutes
	if c < 0 {
		c += minutesPerDay
	}

	return Clock(c)
}

// New creates a Clock from given hour and minute.
func New(hour, minute int) Clock {
	return NewClock(hour*60 + minute)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}

// Add minutes to a Clock.
func (c Clock) Add(minutes int) Clock {
	return NewClock(int(c) + minutes)
}
