// Package clock implements methods to handle clock algorithm.
package clock

import "fmt"

// Clock is struct representing clock using hours and minutes.
type Clock struct {
	minute int
}

// String returns string representation of clock.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minute/60, c.minute%60)
}

// Add returns clock in argument, increased by amount of minutes.
func (c Clock) Add(a int) Clock {
	return New(0, c.minute+a)
}

// Subtract returns clock in argument, decreased by amount of minutes.
func (c Clock) Subtract(a int) Clock {
	return c.Add(-a)
}

// New constructs new Clock from hours and minutes in argument.
func New(h int, m int) Clock {
	m += h * 60
	m %= 24 * 60
	if m < 0 {
		m += 24 * 60
	}
	return Clock{m}
}
