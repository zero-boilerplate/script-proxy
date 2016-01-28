package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"os"
	"strings"
)

var (
	LOGGER *logrus.Entry
)

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

func setupLogs(conf *config, exe string, args []string) *os.File {
	fileHandle, err := os.OpenFile(conf.LogPath, os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Unable to open log file:", err.Error())
	}

	logrus.SetOutput(fileHandle)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)

	formattedArgs := formatArgsWrapWithQuotesIfSpaces(args)
	fullCommandString := fmt.Sprintf("%s %s", exe, strings.Join(formattedArgs, " "))

	LOGGER = logrus.WithField("FULL_CMD", fullCommandString)

	return fileHandle
}
