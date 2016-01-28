package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
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
		logrus.Fatalf("PANIC: %s", errStr)
	}
}
