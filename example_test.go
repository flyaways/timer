package timer_test

import (
	"fmt"
	"time"

	timer "github.com/flyaways/timer"
)

func ExampleTimer() {
	t := timer.New(20, time.Second)
	defer t.Stop()

	done := make(chan struct{})
	go func() {
		//do something
		time.Sleep(time.Second)
		done <- struct{}{}
	}()

	select {
	case <-t.After(time.Second * 5):
		fmt.Println("after", time.Now().Format(time.RFC3339Nano))
	case <-done:
		fmt.Println("done", time.Now().Format(time.RFC3339Nano))
	}
}
