package timer

import (
	"testing"
	"time"

	"github.com/RussellLuo/timingwheel"
	"github.com/TarsCloud/TarsGo/tars/util/rtimer"
)

//BenchmarkTimeWheel benchmarks timewheel.
func BenchmarkFlyawaysTimer(b *testing.B) {
	t := New(20, time.Millisecond*5)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			t.After(time.Millisecond * 100)
		}
	})
}

//BenchmarkTimeWheel benchmarks timewheel.
func BenchmarkTarsGoTimer(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rtimer.After(time.Millisecond * 100)
		}
	})
}

//BenchmarkTimeBase benchmark origin timer.
func BenchmarkOfficalTimer(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			time.After(time.Millisecond * 100)
		}
	})
}

//BenchmarkRussellLuoTimer benchmark timewheel.
func BenchmarkRussellLuoTimer(b *testing.B) {
	tw := timingwheel.NewTimingWheel(time.Millisecond*5, 20)
	tw.Start()
	defer tw.Stop()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			tw.AfterFunc(time.Millisecond*100, func() {})
		}
	})
}
