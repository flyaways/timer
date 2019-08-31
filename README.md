# timer

## Separated from https://github.com/TarsCloud/TarsGo/tree/master/tars/util/rtimer

> for setDeadline and timeout sence ，but performance is better than TarsGo-rtimer

> Thanks a million to https://github.com/TarsCloud/TarsGo/tree/master/tars/util/rtimer

```go
package main

import (
	"github.com/flyaways/timer"
)

func main() {
	//1 step number
	//2 step size
	t := timer.New(20, time.Second)
	t.After(Second * 5)
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
