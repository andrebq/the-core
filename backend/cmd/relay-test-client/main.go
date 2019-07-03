package main

import (
	"bytes"
	"flag"
	"github.com/andrebq/the-core/backend/bridge"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var (
	relay = flag.String("relay", "127.0.0.1:10003", "Relay server")
)

func main() {
	firstClient, err := bridge.DialTCP(*relay)
	if err != nil {
		logEntry().WithError(err).Fatal("unable to establish first client connection")
	}

	secondClient, err := bridge.DialTCP(*relay)
	if err != nil {
		logEntry().WithError(err).Fatal("unable to establish second client connection")
	}

	go ping(firstClient, secondClient.Addr())
	go pong(secondClient)
	select {}
}

func ping(cli *bridge.Client, to bridge.NodeAddress) {
	var pingid uint64
	var sum time.Duration
	var count int64
	for {
		pingid++
		count++
		start := time.Now()
		ping := []byte(time.Now().Format(time.RFC3339Nano))
		err := cli.Send(to, ping)
		if err != nil {
			logEntry().WithError(err).Fatal("unable to ping")
		}
		pong, err := cli.Recv()
		if err != nil {
			logEntry().WithError(err).Fatal("unable to receive pong")
		}
		if pong.From != to {
			logEntry().Error("unexpected pong")
		}
		if !bytes.Equal(ping, pong.Payload) {
			logEntry().Info("invalid pong")
			continue
		}
		delay := time.Now().Sub(start)
		sum += delay
		if count%1000 == 0 {
			avgDelay := time.Duration(int64(sum) / count)
			sum = 0
			count = 0
			logEntry().WithField("ping", pingid).WithField("t", avgDelay).Info("ping-pong")
		}
	}
}

func pong(cli *bridge.Client) {
	for {
		ping, err := cli.Recv()
		if err != nil {
			logEntry().WithError(err).Fatal("unable to recv ping")
		}
		err = cli.Send(ping.From, ping.Payload)
		if err != nil {
			logEntry().WithError(err).Fatal("unable to send pong")
		}
	}
}

func logEntry() *logrus.Entry {
	return logrus.WithField("process", "relay-test-client").
		WithField("pid", os.Getpid())
}
