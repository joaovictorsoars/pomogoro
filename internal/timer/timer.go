package timer

import (
	"context"
	"time"
)

type Timer struct {
	duration  time.Duration
	timerType string
}

func NewTimer(durationSeconds int, timerType string) *Timer {
	return &Timer{
		duration:  time.Duration(durationSeconds) * time.Second,
		timerType: timerType,
	}
}

func (t *Timer) Run(ctx context.Context, updateFn func(mins, secs int, timerType string)) bool {
	endTime := time.Now().Add(t.duration)
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return false
		case currentTime := <-ticker.C:
			if currentTime.After(endTime) {
				return true
			}

			remaining := endTime.Sub(currentTime)
			mins := int(remaining.Minutes())
			secs := int(remaining.Seconds()) % 60

			updateFn(mins, secs, t.timerType)
		}
	}
}
