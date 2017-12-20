package main

import (
	// "fmt"
	"os"
	"strings"
	"path"
	"os/exec"
)

func main() {
	subCommand := strings.Join(os.Args[1:], " ")
	if subCommand == "" {
		subCommand = "/bin/sh"
	}
	script := "source /venv/bin/activate && " + subCommand
	cwd, _ := os.Getwd()
	container := path.Base(cwd) + "_web_1"
	// args := []string{
	// 	"docker",
	// 	"exec",
	// 	"-it",
	// 	container,
	// 	"/bin/sh",
	// 	"-c",
	// 	script,
	// }
	// fmt.Printf("What? %s", args)
	cmd := exec.Command(
		"docker",
		"exec",
		"-it",
		container,
		"/bin/sh",
		"-c",
		script,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Run()
}
