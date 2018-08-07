package main

import (
	"net/http"
	"os"
	"os/signal"
	// "time"

	"github.com/Sirupsen/logrus"

	"github.com/justinrlee/dcos-replication-check/dcos-client"
)

func setupLogger () {
	Formatter := new(logrus.TextFormatter)
	Formatter.TimestampFormat = "2006-01-02 15:04:05"
	Formatter.FullTimestamp = true
	logrus.SetFormatter(Formatter)
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}


func main () {
	setupLogger()
	logrus.Infoln("Status monitor starting.")

	// To detect 
	signalChan := make(chan os.Signal, 1)

	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for _ = range signalChan {
			logrus.Fatalln("Received an interrupt, stopping")
			// ManagerRemoveAllScalers()
			cleanupDone <- true
		}
	}()

	client := dcos-Client.

	router := NewRouter()
	logrus.Fatal(http.ListenAndServe(":8083", router))
}