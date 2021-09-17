package logger

import "log"

// LogError Save error to log
func LogError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
