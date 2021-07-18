package main

import (
	"jekin000/mygoproject/pkg/gologger"
	"fmt"
)


func main() {
	gologger.InitLogger()
	//gologger.Debug(fmt.Sprintf("it is debug,%d",1))
	//gologger.Info(fmt.Sprintf ("it is info ,%d",2))
	gologger.Debug()("%s is pretty.","Tom")
	gologger.Debug()("I think he can do it")
	fmt.Println(gologger.GetLogger().Buf)
	gologger.Debug()("It is ok")
	fmt.Println(gologger.GetLogger().Buf)
}
