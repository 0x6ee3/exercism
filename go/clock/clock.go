package clock

import "fmt"

const testVersion = 4

// Clock type.
type Clock struct {
	hour, minute int
}

// New creates a Clock from given hour and minute. Negative values are supported
func New(hour, minute int) Clock {

	hour += minute / 60

	minute %= 60
	hour %= 24

	return Clock{hour: hour, minute: minute}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add minutes to a Clock. Negative values are supported.
func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}
