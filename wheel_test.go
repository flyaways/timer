package timer

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//TestTimeWheel test timewheel.
func TestTimeWheel(t *testing.T) {
	var deviation int64
	var num int64 = 100
	sleepTime := time.Millisecond * 20000
	timer := New(20, time.Millisecond*5)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			start := time.Now().UnixNano()
			<-timer.After(time.Millisecond * 100)
			end := time.Now().UnixNano()
			d := (end - start) - int64(sleepTime)
			if d >= 0 {
				deviation += d
			} else {
				deviation -= d
			}
			wg.Done()
		}(i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Wait()
	fmt.Println(float64(deviation) / float64(num*int64(sleepTime)))
}
