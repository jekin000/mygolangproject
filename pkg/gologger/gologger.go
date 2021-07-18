package gologger

import (
	"fmt"
	"log"
	//"io"
	"os"
	//"reflect"
	"bytes"
)

type LoggingType int32

const (
	DEBUG  LoggingType = 0
	INFO   LoggingType = 1
	WARN   LoggingType = 2
	ERROR  LoggingType = 3
	ROTATE LoggingType = 4
	UPDATE LoggingType = 5
)

type LogData struct {
	DataType LoggingType
	Data	 string
}
var gLogger *Logger

type Logger struct{
	ioWriter *os.File
	DebugLogger *log.Logger
	infoLogger *log.Logger
	Buf *bytes.Buffer
	Chan chan  *LogData
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
	debugLogger := log.New(logFile, "[Debug]",log.Llongfile|log.LstdFlags)
	infoLogger  := log.New(logFile, "[Info ]",log.Llongfile|log.LstdFlags)

	//fmt.Println(reflect.TypeOf(infoLogger))

	logger := &Logger{
		ioWriter : logFile,
		DebugLogger : debugLogger,
		infoLogger  : infoLogger,
		Buf	    : buf,
		Chan	    : make(chan *LogData),
	}

	gLogger = logger
}

func Close(){
	gLogger.ioWriter.Close()
}

func Debug (str string) {
	logdata  := &LogData {
		DataType : DEBUG,
		Data     : str,
	}
	gLogger.Chan <- logdata
}

func Info (str string){
	gLogger.infoLogger.Println(str)
}

func StartRecorder(){
	go func (){
		for {
			select{
				case logdata :=  <- gLogger.Chan:
					fmt.Println(logdata.Data)
			}
		}
	}()
}
