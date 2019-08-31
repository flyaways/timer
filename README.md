# timer

## Separated from https://github.com/TarsCloud/TarsGo/tree/master/tars/util/rtimer

> for setDeadline and timeout sence ,but performance is better than TarsGo-rtimer.

> Thanks a million to https://github.com/TarsCloud/TarsGo/tree/master/tars/util/rtimer

```go
package main

import (
	"fmt"
	"time"
	"github.com/flyaways/timer"
)

func main() {
	//1 step count：20
	//2 step width：time.Second
	t := timer.New(20, time.Second)
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
```

> Benchmark to FlyawaysTimer, TarsGoTimeWheel and OfficalTimer.


```sh
go test -count=1 -benchtime 3s -benchmem . -bench .

goos: darwin
goarch: amd64
pkg: github.com/flyaways/timer
BenchmarkFlyawaysTimer-4   	200000000	        27.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkTarsGoTimer-4     	 50000000	        78.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkOfficalTimer-4    	 10000000	         396 ns/op	     212 B/op	       3 allocs/op
```
