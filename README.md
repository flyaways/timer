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
```

> Benchmark to FlyawaysTimer, TarsGoTimeWheel and OfficalTimer.


```sh
go test -count=1 -benchtime 5s -benchmem . -bench .
0.9948804265
goos: darwin
goarch: amd64
pkg: github.com/flyaways/timer
cpu: Intel(R) Core(TM) i5-7360U CPU @ 2.30GHz
BenchmarkFlyawaysTimer-4        174450506               53.58 ns/op            0 B/op          0 allocs/op
BenchmarkTarsGoTimer-4          79953462                74.30 ns/op            0 B/op          0 allocs/op
BenchmarkOfficalTimer-4         12491235               519.6 ns/op           201 B/op          3 allocs/op
BenchmarkRussellLuoTimer-4      13017769               514.1 ns/op           164 B/op          2 allocs/op
PASS
ok      github.com/flyaways/timer       43.342s
```
