package clock

import "fmt"

const day = 1440

type Clock struct {
	minute int
}

func New(hour, minute int) Clock {
	total := hour*60 + minute
	total = ((total % day) + day) % day
	return Clock{minute: total}
}

func (c Clock) String() string {
	h := c.minute / 60
	m := c.minute % 60
	return fmt.Sprintf("%02d:%02d", h, m)
}
func (c Clock) Add(m int) Clock {
	return New(0, c.minute+m)
}
func (c Clock) Subtract(m int) Clock {
	return c.Add(-m)
}
