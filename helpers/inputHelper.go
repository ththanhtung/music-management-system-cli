package helpers

import "fmt"

func GetInputInt(msg string) int {
	for {
		fmt.Print(msg)
		var input int
		_, err := fmt.Scanln(&input)

		if err != nil {
			fmt.Println("Invalid input, please enter a valid integer.")
			// Clear the input buffer
			ClearConsole()
			continue
		}
		return input
	}
}

func GetInputString(msg string) string {
	for {
		fmt.Print(msg)
		var input string
		_, err := fmt.Scanln(&input)

		if err != nil {
			fmt.Println("Invalid input, please enter a valid string.")
			// Clear the input buffer
			ClearConsole()
			continue
		}
		return input
	}
}
func GetInputWithDefaultString(msg, defaultValue string) string {
	for {
		fmt.Print(msg)
		var input string
		_, err := fmt.Scanln(&input)

		if err != nil {
			fmt.Println("Invalid input, please enter a valid string.")
			// Clear the input buffer
			ClearConsole()
			continue
		}
		if input == ""{
			return defaultValue
		}
		return input
	}
}
func GetInputWithDefaultInt(msg string, defaultValue int) int {
	for {
		fmt.Print(msg)
		var input int
		_, err := fmt.Scanln(&input)

		if err != nil {
			fmt.Println("Invalid input, please enter a valid string.")
			// Clear the input buffer
			ClearConsole()
			continue
		}
		if input == 0{
			return defaultValue
		}
		return input
	}
}