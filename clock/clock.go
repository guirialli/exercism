package clock

import (
	"fmt"
	"math"
)

func mod(a, b float64) int {
	return int(math.Mod(a, b)+b) % int(b)
}

type Clock struct {
	h int
	m int
}

func New(h, m int) Clock {
	h = mod(float64(h)+float64(m)/60, 24)
	m = mod(float64(m), 60)
	return Clock{
		h: h,
		m: m,
	}
}

func (c Clock) Add(m int) Clock {
	return New(c.h, c.m+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.h, c.m-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}
