package utils

import (
	"log"
	"os"
)

type mylogger struct {
	file string
	*log.Logger
}

var logger *mylogger

func CreateLogger() (logger *mylogger){
	file,_:=os.OpenFile("log",os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	return &mylogger{
		file: "log",
		Logger:log.New(file,"",log.LstdFlags|log.Lshortfile),
	}
}