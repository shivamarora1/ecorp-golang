package main

import (
	"log"
	"os"
)

var (
	highLevelLogger *log.Logger
	lowLevelLogger  *log.Logger
)

func main() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	lowLevelLogger = log.New(file, "", 0)
	lowLevelLogger.Println("File Printing")

	highLevelLogger = log.New(os.Stdout, "", 0)
	highLevelLogger.Println("Terminal Printing", 2, true, 9.089, "Here")
}
