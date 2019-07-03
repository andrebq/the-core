package autoupdate

import (
	"encoding/json"
	get "github.com/hashicorp/go-getter"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
)

type (
	// Application contains the definition of an application and its bundle
	Application struct {
		name     string
		folder   string
		manifest Manifest
	}
)

func Discover(appsURL, appName string) (*Application, error) {
	targetFile := filepath.Join(os.TempDir(), uuid.NewV4().String(), "apps.json")
	defer os.Remove(targetFile)
	content, err := getFileContent(targetFile, appsURL)
	if err != nil {
		return nil, err
	}
	apps := struct {
		Apps []Manifest
	}{}
	err = json.Unmarshal(content, &apps)
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to decode apps file from %v", appsURL)
	}
	for _, m := range apps.Apps {
		if m.Application == appName {
			return NewApplication(m)
		}
	}
	return nil, errors.Errorf("Application %v not found in %v", appName, appsURL)
}

func NewApplication(manifest Manifest) (*Application, error) {
	a := &Application{
		manifest: manifest,
	}
	var err error
	a.folder, err = appFolder(a.manifest.Application)
	return a, err
}

func (a *Application) DownloadBundle(name string) (string, error) {
	os, arch := a.getSameOSArch()
	bundle, ok := a.manifest.getBundle(name, os, arch)
	if !ok {
		return "", errors.Errorf("unable to find bundle %s for [os %s/arch %s].", name, os, arch)
	}
	return a.downloadBundle(bundle)
}

func (a *Application) getSameOSArch() (string, string) {
	return runtime.GOOS, runtime.GOARCH
}

func (a *Application) downloadBundle(b Bundle) (string, error) {
	targetFolder := a.folderForBundle(b)
	downloadURL, err := url.Parse(b.DownloadURL)
	if err != nil {
		return targetFolder, errors.Wrapf(err, "Unable to parse download URL [%s]", b.DownloadURL)
	}
	qs := downloadURL.Query()
	if len(b.Checksum) > 0 {
		qs["checksum"] = []string{b.Checksum}
	}
	downloadURL.RawQuery = qs.Encode()
	err = get.GetAny(targetFolder, downloadURL.String())
	if err != nil {
		return targetFolder, errors.Wrapf(err, "Unable to download to [%v] from [%v]", targetFolder, downloadURL.String())
	}
	return targetFolder, nil
}

func (a *Application) folderForBundle(b Bundle) string {
	switch b.Kind {
	case "bin":
		return binFolder(a.folder)
	case "res":
		return resFolder(a.folder)
	}
	return a.folder
}

func getFileContent(targetFile, inputURL string) ([]byte, error) {
	err := get.GetFile(targetFile, inputURL)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to download %s to %s", inputURL, targetFile)
	}
	data, err := ioutil.ReadFile(targetFile)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to read contents from %s", targetFile)
	}
	return data, nil
}
