package commands

import (
	"focus_timer/services"
	"fmt"
)

// HandleCommand processes user input commands
func HandleCommand(cmd string, timer *services.Timer) {
	switch cmd {
	case "pause":
		timer.Pause()
	case "resume":
		timer.Resume()
	case "reset":
		timer.Reset()
	case "exit":
		timer.Stop()
		fmt.Println("Exiting Focus Timer...")
		return
	default:
		fmt.Println("Invalid command. Use: pause, resume, reset, exit")
	}
}
