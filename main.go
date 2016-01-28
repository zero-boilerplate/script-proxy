package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
)

func handlePanic() {
	if r := recover(); r != nil {
		var errStr string
		switch t := r.(type) {
		case error:
			errStr = t.Error()
			break
		default:
			errStr = fmt.Sprintf("%#v", r)
		}
		log.Fatalf("PANIC: %s", errStr)
	}
}

func getPlatformCommandExeAndArgs(argList []string) (newExe string, newArgs []string) {
	goos := runtime.GOOS
	switch goos {
	case "windows":
		newExe = "cmd"
		newArgs = []string{"/c"}
		newArgs = append(newArgs, argList...)
		return
	case "linux":
		newExe = "bash"
		newArgs = []string{"-c"}
		newArgs = append(newArgs, argList...)
		return
	default:
		panic(fmt.Sprintf("OS '%s' not supported by scroxy", goos))
	}
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

func formatArgsWrapWithQuotesIfSpaces(args []string) []string {
	formattedArgs := args[:]
	for i, a := range formattedArgs {
		if strings.Contains(a, " ") {
			trimmedQuotes := strings.Trim(a, `"`)
			formattedArgs[i] = fmt.Sprintf(`"%s"`, trimmedQuotes)
		}
	}
	return formattedArgs
}

func main() {
	defer handlePanic()

	if len(os.Args) < 2 {
		panic("Scroxy requires at least one command-line argument")
	}

	remainingArgs := os.Args[1:]

	exe, args := getPlatformCommandExeAndArgs(remainingArgs)

	cmd := exec.Command(exe, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	formattedArgs := formatArgsWrapWithQuotesIfSpaces(args)
	fmt.Println(fmt.Sprintf("[SCROXY] running %s %s", exe, strings.Join(formattedArgs, " ")))

	// startTime := time.Now()
	err := cmd.Run()
	// endTime := time.Now()

	exitCode, err := extractExitCodeFromError(err)
	if err != nil {
		panic(err)
	}

	os.Exit(exitCode)
}
