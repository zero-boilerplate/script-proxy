package main

import (
	"os"
	"os/exec"
	"syscall"
)

type stdoutProxy struct{}

func (s *stdoutProxy) Write(b []byte) (int, error) {
	LOGGER.Debug("OUT: " + string(b))
	return os.Stdout.Write(b)
}

type stderrProxy struct{}

func (s *stderrProxy) Write(b []byte) (int, error) {
	LOGGER.Debug("ERR: " + string(b))
	return os.Stderr.Write(b)
}

func extractExitCodeFromError(err error) (int, error) {
	if err == nil {
		return 0, nil
	}

	if exiterr, ok := err.(*exec.ExitError); ok {
		if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
			return status.ExitStatus(), nil
		}
	}

	if err != nil {
		return 1, err
	} else {
		return 0, nil
	}
}

func runCommand(exe string, args []string) int {
	cmd := exec.Command(exe, args...)

	cmd.Stdout = &stdoutProxy{}
	cmd.Stderr = &stderrProxy{}
	cmd.Stdin = os.Stdin

	LOGGER.Infof(`START`)
	err := cmd.Run()
	LOGGER.Infof(`END`)

	exitCode, err := extractExitCodeFromError(err)
	if err != nil {
		panic(err)
	}

	return exitCode
}
