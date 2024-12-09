package helper

import (
	"os/exec"
	"syscall"
)

func Executa(comando string, argumentos ...string) (stdout string, stderr string, exitCode int) {
	cmd := exec.Command(comando, argumentos...)
	combinedOutput, err := cmd.CombinedOutput()
	if err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				return string(combinedOutput), err.Error(), status.ExitStatus()
			}
		}
		return "", err.Error(), -1
	}
	return string(combinedOutput), "", 0
}
