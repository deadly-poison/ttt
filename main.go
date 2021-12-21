package ttt

import (
log "github.com/sirupsen/logrus"
"time"
)

func init() {
	printx()
}

func printx() {
	for {
		time.Sleep(1 * time.Second)
		log.WithFields(log.Fields{
			"animal": "walrus",
		}).Info("A walrus appears")
	}
}
