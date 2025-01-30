package main

import (
	"focus_timer/commands"
	"focus_timer/services"
	"focus_timer/utils"
)

func main() {
	log := utils.GetLogger()
	log.Info("Focus Timer App Started")

	durationInput := utils.GetUserInput("Enter focus duration (minutes): ")
	duration := utils.ParseDuration(durationInput) // Kullanıcı girdisini temizler

	timer := services.NewTimer(duration)
	go timer.Start()

	for {
		cmd := utils.GetUserInput("\nCommand (pause/resume/reset/exit): ")
		commands.HandleCommand(cmd, timer)
		if cmd == "exit" {
			break
		}
	}
}
