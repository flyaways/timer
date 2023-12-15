package timer

import (
	"time"
)

type Timer struct {
	wheel *Wheel
}

// New time range is wheelNum * gearSize = 20*time.Second = 20s
// number：20
// gear：time.Second
func New(wheelNum int, gearSize time.Duration) *Timer {
	if wheelNum <= 0 || gearSize <= 0 {
		return nil
	}

	return &Timer{
		wheel: NewWheel(wheelNum, gearSize),
	}
}

// After like time.After,accuracy means max 1/20 deviation
// the low accuracy the better performance.
func (t *Timer) After(d time.Duration) <-chan struct{} {
	if t.wheel != nil {
		return t.wheel.After(d)
	}

	return t.wheel.After(d)
}

// Stop stops the timewheel
func (t *Timer) Stop() {
	t.wheel.Stop()
}
