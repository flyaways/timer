package timer

import (
	"time"
)

type Timer struct {
	wheel *Wheel
}

func New(count int, step time.Duration) *Timer {
	if count <= 0 || step <= 0 {
		return nil
	}

	return &Timer{
		wheel: NewWheel(count, step),
	}
}

//After like time.After,accuracy means max 1/20 deviation
//low accuracy usually have better performance.
func (t *Timer) After(timeout time.Duration) <-chan struct{} {
	if t.wheel != nil {
		return t.wheel.After(timeout)
	}

	return t.wheel.After(timeout)
}

//Stop stops the timewheel
func (t *Timer) Stop() {
	t.wheel.Stop()
}
