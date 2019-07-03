package main

import (
	"flag"
	"fmt"
	"github.com/andrebq/the-core/backend/autoupdate"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	AppsURL = "https://andrebq.github.io/apps.json"
)

var (
	apps    = flag.String("apps", AppsURL, "URL with the list of available apps")
	appname = flag.String("app", "-", "Application name to update. Must exist in the list")
	bundle  = flag.String("bundle", "bundle", "Bundle to download")
)

func main() {
	flag.Parse()
	app, err := autoupdate.Discover(*apps, *appname)
	if err != nil {
		logEntry().WithError(err).Fatal("unable to discover app")
	}
	target, err := app.DownloadBundle(*bundle)
	if err != nil {
		logEntry().WithError(err).WithField("bundle", *bundle).Fatal("Unable to download bundle")
	}
	fmt.Fprintf(os.Stdout, "%s\n", target)
	return
}

func logEntry() *logrus.Entry {
	return logrus.WithField("process", "autoupdate").
		WithField("apps", *apps).
		WithField("appname", *appname).
		WithField("pid", os.Getpid())
}
