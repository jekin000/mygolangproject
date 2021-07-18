package main

import (
	"jekin000/mygoproject/pkg/gologger"
	"fmt"
)


func main() {
	gologger.InitLogger()
	gologger.StartRecorder()
	//gologger.Debug(fmt.Sprintf("it is debug,%d",1))
	//gologger.Info(fmt.Sprintf ("it is info ,%d",2))
	gologger.Debug(fmt.Sprintf ("%s is busy","Tracy"))
	gologger.Debug("I think he can do it")
	gologger.Debug("It is ok")
}
