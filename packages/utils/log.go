package utils

import (
	"flag"
	log "github.com/sirupsen/logrus"
)

func GetVerboseFlag() *bool {
	return flag.Bool("v", false, "Show more information on run")
}

func TurnOnVerbose(verbose *bool) {
	if *verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}
