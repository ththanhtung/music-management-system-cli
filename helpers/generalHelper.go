package helpers

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

type Helper struct{}

func NewHelper()*Helper{
	return &Helper{}
}

func (h *Helper) ClearInputBuffer() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
}

func (h *Helper) ClearConsole() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (h *Helper) RandomID() string {
	randomID := fmt.Sprint(time.Now().UnixNano())
	return randomID
}