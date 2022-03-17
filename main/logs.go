package main

import (
	"log"
	"os"
)

func GetErrorLog() *log.Logger  {
	file, _ := os.OpenFile("Error.log", os.O_RDWR|os.O_APPEND, 0755)
	return log.New(file, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
}

func GetInfoLog() *log.Logger  {
	file, _ := os.OpenFile("Info.log", os.O_RDWR|os.O_APPEND, 0755)
	return log.New(file, "INFO\t", log.Ldate|log.Ltime)
}