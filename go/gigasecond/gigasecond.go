package gigasecond

import "time"

const testVersion = 4

// AddGigasecond adds 10^9 (1 000 000 000) seconds to a time.
func AddGigasecond(t time.Time) time.Time {
	d, err := time.ParseDuration("1000000000s")
	if err != nil {
		panic("Couldn't create gigasecond duration.")
	}
	return t.Add(d)
}
