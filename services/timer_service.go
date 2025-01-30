package services

import (
	"fmt"
	"time"
	"focus_timer/utils"
)

// Timer struct
type Timer struct {
	totalSeconds int
	pause        bool
	stop         bool
	log          *utils.Logger
}

// NewTimer initializes a new Timer
func NewTimer(minutes int) *Timer {
	return &Timer{
		totalSeconds: minutes * 60,
		pause:        false,
		stop:         false,
		log:          utils.GetLogger(),
	}
}

// Start begins the countdown
func (t *Timer) Start() {
	t.stop = false

	for t.totalSeconds > 0 {
		if t.stop {
			t.log.Info("Timer Stopped and Reset")
			return
		}

		if t.pause {
			time.Sleep(1 * time.Second)
			continue
		}

		mins := t.totalSeconds / 60
		secs := t.totalSeconds % 60
		fmt.Printf("\rTime Left: %02d:%02d", mins, secs)

		time.Sleep(1 * time.Second)
		t.totalSeconds--
	}

	t.log.Info("Focus Session Completed!")
	fmt.Println("\nFocus Session Completed! Take a break.")
}

// Pause suspends the timer
func (t *Timer) Pause() {
	if !t.pause {
		t.pause = true
		t.log.Info("Timer Paused")
		fmt.Println("\nTimer Paused")
	} else {
		fmt.Println("Timer is already paused!")
	}
}

// Resume continues the timer
func (t *Timer) Resume() {
	if t.pause {
		t.pause = false
		t.log.Info("Timer Resumed")
		fmt.Println("\nTimer Resumed")
	} else {
		fmt.Println("Timer is not paused!")
	}
}

// Reset stops and restarts the timer
func (t *Timer) Reset() {
	t.stop = true
	fmt.Println("\nTimer Reset! Restarting...")
	go t.Start()
}

// Stop terminates the timer
func (t *Timer) Stop() {
	t.stop = true
	t.log.Info("Timer Stopped")
}
