package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// GetUserInput reads user input from terminal
func GetUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(strings.ToLower(input))
}
