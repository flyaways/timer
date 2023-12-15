package timer_test

import (
	"testing"
	"time"

	"github.com/RussellLuo/timingwheel"
	"github.com/TarsCloud/TarsGo/tars/util/rtimer"
	timer "github.com/flyaways/timer"
)

const (
	gearSize      time.Duration = time.Millisecond * 5
	afterDuration time.Duration = time.Millisecond * 100
	wheelNum      int           = 20
)

// BenchmarkTimeWheel benchmarks timewheel.
func BenchmarkFlyawaysTimer(b *testing.B) {
	t := timer.New(wheelNum, gearSize)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			t.After(afterDuration)
		}
	})
}

// BenchmarkTimeWheel benchmarks timewheel.
func BenchmarkTarsGoTimer(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rtimer.After(afterDuration)
		}
	})
}

// BenchmarkTimeBase benchmark origin timer.
func BenchmarkOfficalTimer(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			time.After(afterDuration)
		}
	})
}

// BenchmarkRussellLuoTimer benchmark timewheel.
func BenchmarkRussellLuoTimer(b *testing.B) {
	tw := timingwheel.NewTimingWheel(gearSize, int64(wheelNum))
	tw.Start()
	defer tw.Stop()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			tw.AfterFunc(afterDuration, func() {})
		}
	})
}
