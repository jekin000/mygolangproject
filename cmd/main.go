package main

import (
	"log"
	"os"
)

func main() {
	fileName := "/var/log/mydebug.log"
	logFile, err := os.Create(fileName)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("Open file failed")
	}
	debugLog := log.New(logFile, "[Debug]", log.Llongfile)
	debugLog.Println("A debug msg here")
	debugLog.SetPrefix("[Info]")
	debugLog.Println("A info msg here")
	debugLog.SetFlags(debugLog.Flags() | log.LstdFlags)
	debugLog.Println("A different prefix")
}
