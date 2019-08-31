# timer

## Separated from https://github.com/TarsCloud/TarsGo/tree/master/tars/util/rtimer

> for quic learn and test,and maybe modify it later.

> Thanks a million to https://github.com/TarsCloud/TarsGo/tree/master/tars/util/rtimer

```go
package main

import (
	"fmt"

	"github.com/flyaways/timer"
)

func main() {
	t := timer.New(20, time.Second)
	t.After(Second * 5)
}
```
