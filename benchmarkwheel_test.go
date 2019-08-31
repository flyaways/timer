package timer

import (
	"testing"
	"time"

	"github.com/TarsCloud/TarsGo/tars/util/rtimer"
)

//BenchmarkTimeWheel benchmarks timewheel.
func BenchmarkFlyawaysTimer(b *testing.B) {
	t := New(20, time.Millisecond*5)
	for i := 0; i < b.N; i++ {
		t.After(time.Millisecond * 100)
	}
}

//BenchmarkTimeWheel benchmarks timewheel.
func BenchmarkTarsGoTimer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rtimer.After(time.Millisecond * 100)
	}
}

//BenchmarkTimeBase benchmark origin timer.
func BenchmarkOfficalTimer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.After(time.Millisecond * 100)
	}
}
