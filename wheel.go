package timer

import (
	"sync"
	"time"
)

type Wheel struct {
	ticker   *time.Ticker
	wheels   []chan struct{}
	gearSize time.Duration
	wheelNum int
	current  int
	lock     sync.RWMutex
}

// NewWheel new timewheel with gearSize and wheelNum.
func NewWheel(wheelNum int, gearSize time.Duration) *Wheel {
	w := &Wheel{
		gearSize: gearSize,
		wheelNum: wheelNum + 1,
		wheels:   make([]chan struct{}, wheelNum+1),
	}

	for i := 0; i < wheelNum+1; i++ {
		w.wheels[i] = make(chan struct{})
	}

	w.ticker = time.NewTicker(gearSize)

	go w.run()

	return w
}

// Stop the timewheel
func (w *Wheel) Stop() {
	w.ticker.Stop()
}

// After like time.After
func (w *Wheel) After(timeout time.Duration) <-chan struct{} {
	w.lock.RLock()
	index := (int(timeout/w.gearSize) + w.current) % w.wheelNum
	c := w.wheels[index]
	w.lock.RUnlock()

	return c
}

func (w *Wheel) run() {
	for range w.ticker.C {
		w.lock.Lock()
		oldestC := w.wheels[w.current]
		w.wheels[w.current] = make(chan struct{})
		w.current = (w.current + 1) % w.wheelNum
		w.lock.Unlock()

		close(oldestC)
	}
}
