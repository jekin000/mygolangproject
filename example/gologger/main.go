package main

import (
	"jekin000/mygoproject/pkg/gologger"
	"fmt"
	"time"
	"os"
	"runtime"
)

func dolog(num int){
	for i:=0;i<30;i++ {
		gologger.Debug(fmt.Sprintf ("%s:[%d]: run %d loop",runFuncName(),num,i))
		time.Sleep(time.Duration(1)*time.Second)
	}
}

func  runFuncName() string{
	pc := make([]uintptr,1)
	runtime.Callers(2,pc)
    	f := runtime.FuncForPC(pc[0])
    	return f.Name()
}

func main() {
	gologger.InitLogger()
	gologger.StartRecorder()
	for i:=0;i<10;i++{
		go dolog(i)
	}

	time.Sleep(time.Duration(5)*time.Second)
	os.Rename("/var/log/mydebug.log","/var/log/mydebug.log.1")
	os.Create("/var/log/mydebug.log")
	gologger.LogRotate()

	time.Sleep(time.Duration(70)*time.Second)
}
