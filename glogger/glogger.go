package glogger

import (
	"log"
	"os"
	"sync"
)

type gLogger struct {
	*log.Logger
	filename string
}

var glogger *gotalkLogger
var once sync.Once

//GetInstance create a singleton instance of the glogger
func GetInstance() *gotalkLogger {
	once.Do(func() {
		glogger = createLogger("gotalkLogger.log")
	})
	return glogger
}

//Create a logger instance
func createLogger(fname string) *gotalkLogger {
	file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

	return &gotalkLogger{
		filename: fname,
		Logger:   log.New(file, "gotalk ", log.Lshortfile),
	}
}
