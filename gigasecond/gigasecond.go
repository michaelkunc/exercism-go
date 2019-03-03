// Package gigasecond returns a gigasecond for a given time
package gigasecond

import "time"

// AddGigasecond returns a timestamp afer 10^9 seconds
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}
