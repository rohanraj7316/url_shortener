package helpers

import (
	"log"
	"os"
)

var (
	// Error used to log error
	Error *log.Logger
	// Info used to log info
	Info *log.Logger
	// Fatal used to log error which have serious implications
	Fatal *log.Logger
)

// InitLogger initalizing logger
func InitLogger() {
	file, err := os.OpenFile("server.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	Error = log.New(file, "[ ERROR ] ", log.LstdFlags|log.Llongfile)
	Info = log.New(file, "[ INFO ] ", log.LstdFlags|log.Llongfile)
	Fatal = log.New(file, "[ FATAL ]", log.LstdFlags|log.Llongfile)
}
