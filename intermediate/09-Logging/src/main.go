package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("This is an info message")
	log.Warn("This is a warning message")
	log.Error("This is an error message")

	// Logging with fields (structured logging)
	log.WithFields(log.Fields{
		"user_id":   123,
		"user_name": "John",
	}).Info("User logged in")
}
