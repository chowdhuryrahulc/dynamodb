package logger

import "log"

func PANIC(message string, err error) {
	// this shows panic error
	if err != nil {
		log.Panic(message, err)
	}
}

func INFO(message string, data interface{}) {
	// this prints out the info
	log.Print(message, data)
}
