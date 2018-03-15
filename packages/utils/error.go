package utils

import (
	log "github.com/sirupsen/logrus"
)

func HandleUnrecoverableError(err error) {
	if err != nil {
		log.Errorf("Unrecoverable error appeard: %s", err.Error())
		log.Panic("Application failed ... Goodbye!")
	}
}
