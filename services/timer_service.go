package services

import (
	"fmt"
	"time"
	"focus_timer/utils"
)

// Timer struct to manage countdown
type Timer struct {
	TotalSeconds int
	Pause        bool
	Stop         bool
	Log          *utils.Logger
}

// NewTimer initializes a new Timer
func NewTimer(minutes int) *Timer {
	return &Timer{
		TotalSeconds: minutes * 60,
		Pause:        false,
		Stop:         false,
		Log:          utils.GetLogger(),
	}
}

// Start begins the countdown timer
func (t *Timer) Start() {
	t.Stop = false
	t.Log.Info("Timer started for", t.TotalSeconds/60, "minutes.")

	for t.TotalSeconds > 0 {
		// Stop the timer if reset
		if t.Stop {
			t.Log.Warn("Timer Stopped and Reset")
			return
		}

		// Pause the timer if requested
		if t.Pause {
			t.Log.Warn("Timer Paused")
			time.Sleep(1 * time.Second)
			continue
		}

		// Update the timer display
		mins := t.TotalSeconds / 60
		secs := t.TotalSeconds % 60
		t.Log.Info(fmt.Sprintf("Time Left: %02d:%02d", mins, secs))

		time.Sleep(1 * time.Second)
		t.TotalSeconds--
	}

	t.Log.Info("Focus Session Completed! Take a break.")
}
