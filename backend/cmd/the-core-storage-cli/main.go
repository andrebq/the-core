package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"github.com/andrebq/the-core/backend/api"
	"github.com/sirupsen/logrus"
	grpc "google.golang.org/grpc"
	"io/ioutil"
	"os"
	"strings"
)

var (
	server = flag.String("server", "localhost:20001", "Blob server address")
	get    = flag.Bool("get", false, "Read ContentIDs from stdin and print them to stdout, logs any errors and abort")
)

func main() {
	flag.Parse()
	cliconn, err := grpc.Dial(*server, grpc.WithInsecure(), grpc.WithUserAgent("the-core-storage-cli"))
	if err != nil {
		logEntry().WithError(err).WithField("server", *server).Fatal("unable to dial")
	}

	storage := api.NewStorageClient(cliconn)
	if *get {
		doGet(storage)
	}
	doPut(storage)
}

func doGet(conn api.StorageClient) {
	sc := bufio.NewScanner(os.Stdin)
	sz := int64(0)
	for sc.Scan() {
		id := sc.Text()
		parts := strings.Split(id, ":")
		if len(parts) != 2 {
			logEntry().WithError(fmt.Errorf("invalid id: %v", id)).Fatal()
		}
		content, err := conn.Get(context.Background(), &api.ContentID{
			Alg:     parts[0],
			HexHash: parts[1],
		})
		if err != nil {
			logEntry().WithError(err).WithField("contentID", id).Fatal("unable to read content")
			continue
		}
		n, _ := os.Stdout.Write(content.Body)
		sz += int64(n)
		logEntry().WithField("contentID", id).WithField("length", len(content.Body)).WithField("written", sz).Info()
	}
}

func doPut(conn api.StorageClient) {
	data := readAllInput()
	id, err := conn.Put(context.Background(), &api.Content{
		Body: data,
	})
	if err != nil {
		logEntry().WithError(err).WithField("length", len(data)).Fatal("unable to send data")
	}
	fmt.Fprintf(os.Stdout, "%s:%s", id.Alg, id.HexHash)
}

func logEntry() *logrus.Entry {
	return logrus.WithField("process", "the-core-storage-cli").
		WithField("pid", os.Getpid())
}

func readAllInput() []byte {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		logEntry().WithError(err).Fatal("unable to read stdin")
	}
	return data
}
