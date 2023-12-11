package logger

import "log"

//log panic
func PANIC(message string, err error) {
	if err != nil {
		//return log panic
		log.Panic(message, err)
	}
}

//log info
func INFO(message string, data interface{}) {
	log.Print(message, data)
}
