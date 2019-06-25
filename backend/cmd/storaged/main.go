package main

import (
	"flag"
	"os"
	"github.com/sirupsen/logrus"
	grpc "google.golang.org/grpc"
	"net"
	"github.com/andrebq/the-core/backend/storage"
	"github.com/andrebq/the-core/backend/storage/dskdrv"
	"github.com/andrebq/the-core/backend/server"
	"github.com/andrebq/the-core/backend/api"
)

var (
	bind = flag.String("bind", "0.0.0.0:20001", "Address to bind and listen for incoming requests")
	basePath = flag.String("base", "storaged_data", "Folder to store files")
)

func main() {
	flag.Parse()
	blobStorage, err := storage.New(storage.WithDriver(dskdrv.New(*basePath)))
	if err != nil {
		logEntry().WithError(err).
			WithField("basePath", *basePath).Fatal("unable to open blob storage")
	}


	lst, err := net.Listen("tcp4", *bind)
	if err != nil {
		logEntry().WithError(err).WithField("bind", *bind).Fatal("unable to bind")
	}

	grpcServer := grpc.NewServer()
	storageSrv := server.NewStorage(blobStorage)
	api.RegisterStorageServer(grpcServer, storageSrv)

	logEntry().WithField("bind", *bind).WithField("basePath", *basePath).Info("starting...")
	err = grpcServer.Serve(lst)
	if err != nil {
		logEntry().WithError(err).Fatal("unable to start grpc")
	}
	logEntry().Info("graceful shutdown")
}

func logEntry() *logrus.Entry {
	return logrus.WithField("process", "storaged").
		WithField("pid", os.Getpid())
}