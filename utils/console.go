package utils

import (
	"os"
	"os/exec"
	"runtime"
)

/*
Console stores util methods for terminal
*/
type Console struct{}

/*
Clear clear the screen
*/
func (m *Console) Clear() {
	switch ops := runtime.GOOS; ops {
	case "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
