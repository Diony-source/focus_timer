package main

import (
	"focus_timer/commands"
	"focus_timer/services"
	"focus_timer/utils"
	"sync"
)

func main() {
	log := utils.GetLogger()
	log.Info("Focus Timer App Started")

	// Get user input for focus session duration
	durationInput := utils.GetUserInput("Enter focus duration (minutes): ")
	duration := utils.ParseDuration(durationInput)

	// Initialize and start the timer
	timer := services.NewTimer(duration)

	var wg sync.WaitGroup
	wg.Add(1)

	// Start the timer in a separate goroutine
	go func() {
		timer.Start()
		wg.Done()
	}()

	// Handle user commands separately
	for {
		cmd := utils.GetUserInput("\n\nCommand (pause/resume/reset/exit): ")
		log.Info("User entered command:", cmd)
		commands.HandleCommand(cmd, timer)
		if cmd == "exit" {
			log.Info("User exited the Focus Timer.")
			break
		}
	}

	wg.Wait()
}
