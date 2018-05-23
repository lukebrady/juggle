package logging

import (
	"log"
	"os"
)

func NewLogger(path string) error {
	logger, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	log.SetOutput(logger)
	return nil
}
