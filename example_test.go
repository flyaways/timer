package timer

import (
	"fmt"
	"time"
)

func ExampleTimer() {
	t := New(20, time.Second)
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
