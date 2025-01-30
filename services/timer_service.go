package services

import (
	"fmt"
	"time"
	"focus_timer/utils"
)

// Timer struct to manage countdown
type Timer struct {
	TotalSeconds int
	isPaused     bool
	isStopped    bool
	Log          *utils.Logger
}

// NewTimer initializes a new Timer
func NewTimer(minutes int) *Timer {
	return &Timer{
		TotalSeconds: minutes * 60,
		isPaused:     false,
		isStopped:    false,
		Log:          utils.GetLogger(),
	}
}

// Start begins the countdown timer
func (t *Timer) Start() {
	t.isStopped = false
	t.Log.Info("Timer started for", t.TotalSeconds/60, "minutes.")

	for t.TotalSeconds > 0 {
		if t.isStopped {
			t.Log.Warn("Timer Stopped and Reset")
			return
		}

		if t.isPaused {
			t.Log.Warn("Timer Paused")
			time.Sleep(1 * time.Second)
			continue
		}

		// Update the timer display
		mins := t.TotalSeconds / 60
		secs := t.TotalSeconds % 60
		fmt.Printf("\rTime Left: %02d:%02d  ", mins, secs)

		time.Sleep(1 * time.Second)
		t.TotalSeconds--
	}

	t.Log.Info("Focus Session Completed! Take a break.")
}

// Pause pauses the timer
func (t *Timer) Pause() {
	if !t.isPaused {
		t.isPaused = true
		t.Log.Warn("Timer Paused")
	} else {
		t.Log.Warn("Timer is already paused")
	}
}

// Resume resumes the timer
func (t *Timer) Resume() {
	if t.isPaused {
		t.isPaused = false
		t.Log.Info("Timer Resumed")
	} else {
		t.Log.Warn("Timer is not paused")
	}
}

// Reset stops the timer and resets it
func (t *Timer) Reset() {
	t.isStopped = true
	t.TotalSeconds = 0
	t.Log.Info("Timer Reset")
}

// Stop completely stops the timer
func (t *Timer) Stop() {
	t.isStopped = true
	t.Log.Info("Timer Stopped")
}
