package timer

import (
	"time"
)

type Timer struct {
	wheel *Wheel
}

// New time range is number * gear = 20*time.Second = 20s
// number：20
// gear：time.Second
func New(number int, gear time.Duration) *Timer {
	if number <= 0 || gear <= 0 {
		return nil
	}

	return &Timer{
		wheel: NewWheel(number, gear),
	}
}

// NewTimer time range is number * gear = 60000*time.Millisecond = 60s
func NewTimer() *Timer {
	return &Timer{
		wheel: NewWheel(60000, time.Millisecond),
	}
}

// After like time.After,accuracy means max 1/20 deviation
// low accuracy usually have better performance.
func (t *Timer) After(timeout time.Duration) <-chan struct{} {
	if t.wheel != nil {
		return t.wheel.After(timeout)
	}

	return t.wheel.After(timeout)
}

// Stop stops the timewheel
func (t *Timer) Stop() {
	t.wheel.Stop()
}
