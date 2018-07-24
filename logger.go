package main

import (
	"log"
	"os"
)

// NewLogger creates a logger that can be used throughout juggle's execution.
func NewLogger(path string) error {
	// Open the log file that will be used to log events.
	logger, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	log.SetOutput(logger)
	return nil
}
