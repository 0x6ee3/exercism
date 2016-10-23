package clock

import "fmt"

const testVersion = 4

// Clock type.
type Clock struct {
	hour, minute int
}

// New creates a Clock from given hour and minute. Negative values are supported
func New(hour, minute int) Clock {

	// Add or substract hours if minutes "overflow"
	// works because minute and hour are int
	hour += minute / 60

	minute %= 60
	hour %= 24

	// Fix negative values
	if minute < 0 {
		minute = 60 + minute
		hour--
	}
	if hour < 0 {
		hour = 24 + hour
	}

	return Clock{hour: hour, minute: minute}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add minutes to a Clock. Negative values are supported.
func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}
