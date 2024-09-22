package gigasecond

import "time"

// Just multiply 1*10^9 for the constant time.Second (1000)
const gigaSecond = 1e9 * time.Second

// This function recived a time and return time + gigaSecond
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigaSecond)

}
