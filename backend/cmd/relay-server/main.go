package main

import (
	"flag"
	"github.com/andrebq/the-core/backend/bridge"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	bind = flag.String("bind", "0.0.0.0:10003", "Address to listen for relay clients")
)

func main() {
	flag.Parse()

	relay := bridge.NewRelay()
	logEntry().WithField("bind", *bind).Info("starting relay...")
	err := relay.RunTCP(*bind)
	if err != nil {
		logEntry().WithError(err).Error("error running tcp relay")
	}
}

func logEntry() *logrus.Entry {
	return logrus.WithField("process", "relay-server").
		WithField("pid", os.Getpid())
}
