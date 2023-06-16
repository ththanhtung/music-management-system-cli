package helpers

import (
	"fmt"
)

func (h *Helper) GetInputInt(msg string) int {
	for {
		fmt.Print(msg)
		var input int
		_, err := fmt.Scanln(&input)

		if err != nil {
			fmt.Println("Invalid input, please enter a valid integer.")
			// Clear the input buffer
			h.ClearConsole()
			continue
		}
		return input
	}
}

func (h *Helper) GetInputString(msg string) string {

	fmt.Print(msg)
	var input string
	fmt.Scanln(&input)
	return input

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
		if input == 0 {
			return defaultValue
		}
		return input
	}
}
