package main

import (
	"os"
	"fmt"
)

const LOG = "/home/louis/golog"

func errorCheck(err error) {
// simple error checker

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func log(msg string, a ...interface{}) {
// format then append string to the log file

	f, err := os.OpenFile(LOG, os.O_APPEND|os.O_WRONLY, 0600)
	errorCheck(err)
	_, err = f.WriteString(fmt.Sprintf(msg, a...) + "\n")
	errorCheck(err)
}
