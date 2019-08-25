package main

import (
	"os"
	"fmt"
)

const LOG = "/home/louis/golog"

// format then append string to log file
func log(msg string, a ...interface{}) {
	f, err := os.OpenFile(LOG, os.O_APPEND|os.O_WRONLY, 0600)
	errorCheck(err)
	_, err = f.WriteString(fmt.Sprintf(msg, a...) + "\n")
	errorCheck(err)
}
