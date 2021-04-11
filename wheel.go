package timer

import (
	"sync"
	"time"
)

type Wheel struct {
	ticker  *time.Ticker
	wheels  []chan struct{}
	gear    time.Duration
	number  int
	current int
	lock    sync.RWMutex
}

//NewWheel new timewheel with gear and number.
func NewWheel(number int, gear time.Duration) *Wheel {
	w := &Wheel{
		gear:   gear,
		number: number + 1,
		wheels: make([]chan struct{}, number+1),
	}

	for i := 0; i < number+1; i++ {
		w.wheels[i] = make(chan struct{})
	}

	w.ticker = time.NewTicker(gear)

	go w.run()

	return w
}

//Stop the timewheel
func (w *Wheel) Stop() {
	w.ticker.Stop()
}

//After like time.After
func (w *Wheel) After(timeout time.Duration) <-chan struct{} {
	w.lock.RLock()
	index := (int(timeout/w.gear) + w.current) % w.number
	c := w.wheels[index]
	w.lock.RUnlock()

	return c
}

func (w *Wheel) run() {
	for range w.ticker.C {
		w.lock.Lock()
		oldestC := w.wheels[w.current]
		w.wheels[w.current] = make(chan struct{})
		w.current = (w.current + 1) % w.number
		w.lock.Unlock()

		close(oldestC)
	}
}
