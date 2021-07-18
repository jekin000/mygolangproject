package gologger

import (
	//"fmt"
	"log"
	//"io"
	"os"
	//"reflect"
	"bytes"
)

var gLogger *Logger

type Logger struct{
	ioWriter *os.File
	DebugLogger *log.Logger
	infoLogger *log.Logger
	Buf *bytes.Buffer
}

func GetLogger() *Logger {
	return gLogger
}

func InitLogger() {
	fileName := "/var/log/mydebug.log"
	logFile, err := os.OpenFile(fileName,os.O_APPEND|os.O_WRONLY,os.ModeAppend)
	if err != nil {
		log.Fatalln("Open file failed,%s",err)
	}

	buf         := bytes.NewBufferString("")
	debugLogger := log.New(buf, "[Debug]",log.Llongfile|log.LstdFlags)
	infoLogger  := log.New(logFile, "[Info ]",log.Llongfile|log.LstdFlags)

	//fmt.Println(reflect.TypeOf(infoLogger))

	logger := &Logger{
		ioWriter : logFile,
		DebugLogger : debugLogger,
		infoLogger  : infoLogger,
		Buf	    : buf,
	}

	gLogger = logger
}

func Close(){
	gLogger.ioWriter.Close()
}

func Debug () func (format string,v ...interface{}){
	return gLogger.DebugLogger.Printf
}

func Info (str string){
	gLogger.infoLogger.Println(str)
}
