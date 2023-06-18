package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (h *Helper) GetInputInt(msg string) int {
	for {
		fmt.Print(msg)
		var input int
		_, err := fmt.Scanln(&input)

		if err != nil {
			// Clear the input buffer
			h.ClearConsole()
			fmt.Println("Invalid input, please enter a valid integer.")
			continue
		}
		return input
	}
}

func (h *Helper) GetInputString(msg string) string {
	fmt.Print(msg)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error occurred while reading input:", err)
		return ""
	}

	// Remove the newline character at the end of the input string
	input = strings.TrimSuffix(input, "\n")
	return strings.TrimSpace(input)
}
func (h *Helper) GetInputWithDefaultString(msg, defaultValue string) string {
	input := h.GetInputString(msg)

	if input == "" {
		return defaultValue
	}
	return input
}

func (h *Helper) GetInputWithDefaultInt(msg string, defaultValue int) int {
	for {
		fmt.Print(msg)
		var input int
		_, err := fmt.Scanln(&input)

		if err != nil {
			fmt.Println("Invalid input, please enter a valid string.")
			// Clear the input buffer
			h.ClearConsole()
			continue
		}
		if input == -1 {
			return defaultValue
		}
		return input
	}
}

func (h *Helper) GetInputIntWithConstraint(msg string, defaultValue, lowerBound, upperBound int) int {
	for {
		input := h.GetInputInt(msg)
		if input == -1 {
			return defaultValue
		}
		if input < lowerBound {
			fmt.Println("Input cannot less than", lowerBound)
			continue
		}
		if input > upperBound {
			fmt.Println("Input cannot greater than", upperBound)
			continue
		}
		return input
	}
}

func (h *Helper) GetInputStringWithConstraint(msg, defaultValue string, lowerBound, upperBound int) string {
	for {
		input := h.GetInputString(msg)
		if input == "" {
			return defaultValue
		}
		if len(input) < lowerBound {
			fmt.Println("Input cannot less than", lowerBound,"characters")
			continue
		}
		if len(input) > upperBound {
			fmt.Println("Input cannot greater than", upperBound,"character")
			continue
		}
		return input
	}
}
