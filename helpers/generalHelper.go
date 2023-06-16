package helpers

import (
	"bufio"
	"os"
	"os/exec"
	"runtime"
)

func ClearInputBuffer() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
}

func ClearConsole() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
