package main

import (
	"os"
)

func main() {
	defer handlePanic()

	if len(os.Args) < 2 {
		panic("Scroxy requires at least one command-line argument")
	}

	loadFullExePath()

	requestedArgs := os.Args[1:]
	exe, args := getPlatformCommandExeAndArgs(requestedArgs)

	conf := loadConfig()
	fileHandle := setupLogs(conf, exe, args)
	defer fileHandle.Close()

	exitCode := runCommand(exe, args)
	os.Exit(exitCode)
}
