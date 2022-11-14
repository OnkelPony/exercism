/*
Package gigasecond implements the way to count, how much time is spent within one gigaseeond
 */
package gigasecond

// import path for the time package from the standard library
import "time"

const durationInSeconds = 1000000000

// returns time passed during one gigasecond, which is added to the time in parameter
func AddGigasecond(t time.Time) time.Time {

	afterOneGigasecond := t.Add(time.Second * durationInSeconds)
	return afterOneGigasecond
}
