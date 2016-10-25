package gigasecond

import "time"

const (
	testVersion = 4

	// Gigasecond is 10^9 seconds
	Gigasecond = 1000000000
)

// AddGigasecond adds 10^9 (1 000 000 000) seconds to a time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * Gigasecond)
}
