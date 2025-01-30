package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GetUserInput reads user input from terminal
func GetUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(strings.ToLower(input))
}

// ParseDuration converts user input into integer minutes
func ParseDuration(input string) int {
	duration, err := strconv.Atoi(input)
	if err != nil || duration <= 0 {
		fmt.Println("Invalid input! Defaulting to 25 minutes.")
		return 25
	}
	return duration
}
