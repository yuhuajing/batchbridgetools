package utils

import (
	"fmt"
	"log"
	"main/config"
	"os"
)

func Intolog(msg string) {
	fmt.Println(msg)
	// open log file
	logFile, err := os.OpenFile(config.LogConfig.LogFile, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	// Set log out put and enjoy :)
	log.SetOutput(logFile)

	// optional: log date-time, filename, and line number
	//log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Print(msg)

}
