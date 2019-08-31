# timer

## Separated from https://github.com/TarsCloud/TarsGo/tree/master/tars/util/rtimer

> for setDeadline and timeout.

> Thanks a million to https://github.com/TarsCloud/TarsGo/tree/master/tars/util/rtimer

```go
package main

import (
	"github.com/flyaways/timer"
)

func main() {
	t := timer.New(20, time.Second)
	t.After(Second * 5)
}
```
