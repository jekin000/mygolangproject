package main

import (
	"jekin000/mygoproject/pkg/gologger"
	"fmt"
)


func main() {
	logger := gologger.InitLogger()
	logger.Debug(fmt.Sprintf("it is debug,%d",1))
	logger.Info(fmt.Sprintf ("it is info ,%d",2))
}
