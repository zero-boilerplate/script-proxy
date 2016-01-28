package main

import (
	"fmt"
	"runtime"
)

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
