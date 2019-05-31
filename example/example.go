package main

import (
	boltlogrus "github.com/kennykarnama/logrus-bolt-hook"
	log "github.com/sirupsen/logrus"
)

func main() {
	boltHook := boltlogrus.NewBoltHook(boltlogrus.Dbpath("customlog.db"))
	log.AddHook(boltHook)
	log.WithFields(log.Fields{
		"animal": "Dog",
		"number": "1",
	}).Info("Kenny")

}
