package timer

import (
	"sync"
	"time"
)

type Wheel struct {
	ticker  *time.Ticker
	wheels  []chan struct{}
	step    time.Duration
	count   int
	current int
	lock    sync.RWMutex
}

//NewWheel new timewheel with step and count.
func NewWheel(count int, step time.Duration) *Wheel {
	w := &Wheel{
		step:   step,
		count:  count + 1,
		wheels: make([]chan struct{}, count+1),
	}

	for i := 0; i < count+1; i++ {
		w.wheels[i] = make(chan struct{})
	}

	w.ticker = time.NewTicker(step)

	go w.run()

	return w
}

//Stop the timewheel
func (w *Wheel) Stop() {
	w.ticker.Stop()
}

//After like time.After
func (w *Wheel) After(timeout time.Duration) <-chan struct{} {
	index := (int(timeout/w.step) + w.current) % w.count
	w.lock.RLock()
	c := w.wheels[index]
	w.lock.RUnlock()
	return c
}

func (w *Wheel) run() {
	for range w.ticker.C {
		w.lock.Lock()
		oldestC := w.wheels[w.current]
		w.wheels[w.current] = make(chan struct{})
		w.current = (w.current + 1) % w.count
		w.lock.Unlock()
		close(oldestC)
	}
}
