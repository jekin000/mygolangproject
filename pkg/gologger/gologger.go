package gologger

import (
	//"fmt"
	"log"
	//"io"
	"os"
	//"reflect"
)

type Logger struct{
	ioWriter *os.File
	debugLogger *log.Logger
	infoLogger *log.Logger
}

func InitLogger() *Logger{
	fileName := "/var/log/mydebug.log"
	logFile, err := os.OpenFile(fileName,os.O_APPEND|os.O_WRONLY,os.ModeAppend)
	if err != nil {
		log.Fatalln("Open file failed,%s",err)
	}
	debugLogger := log.New(logFile, "[Debug]",log.Llongfile|log.LstdFlags)
	infoLogger  := log.New(logFile, "[Info ]",log.Llongfile|log.LstdFlags)

	//fmt.Println(reflect.TypeOf(infoLogger))

	logger := &Logger{
		ioWriter : logFile,
		debugLogger : debugLogger,
		infoLogger  : infoLogger,
	}

	return logger
}

func (l *Logger) Close(){
	l.ioWriter.Close()
}

func (l *Logger) Debug (str string){
	l.debugLogger.Println(str)
}

func (l *Logger) Info (str string){
	l.infoLogger.Println(str)
}
